package tf

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/chanzuckerberg/happy/shared/config"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/ext/typeexpr"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/hashicorp/terraform-config-inspect/tfconfig"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/gocty"
)

type provider struct {
	Name    string
	Source  string
	Version string
}

type variable struct {
	Name        string
	Type        string
	Description string
	Default     cty.Value
}

const autogeneratedComment = "# This file is autogenerated by 'happy infra generate'. Do not edit manually."

var requiredProviders []provider = []provider{
	{
		Name:    "aws",
		Source:  "hashicorp/aws",
		Version: ">= 4.65",
	},
	{
		Name:    "kubernetes",
		Source:  "hashicorp/kubernetes",
		Version: ">= 2.20",
	},
	{
		Name:    "datadog",
		Source:  "datadog/datadog",
		Version: ">= 3.20.0",
	},
	{
		Name:    "happy",
		Source:  "chanzuckerberg/happy",
		Version: ">= 0.53.5",
	},
}

var requiredVariables []variable = []variable{
	{
		Name:        "aws_account_id",
		Type:        "string",
		Description: "AWS account ID to apply changes to",
	},
	{
		Name:        "k8s_cluster_id",
		Type:        "string",
		Description: "EKS K8S Cluster ID",
	},
	{
		Name:        "k8s_namespace",
		Type:        "string",
		Description: "K8S namespace for this stack",
	},
	{
		Name:        "aws_role",
		Type:        "string",
		Description: "Name of the AWS role to assume to apply changes",
	},
	{
		Name:        "image_tag",
		Type:        "string",
		Description: "Please provide an image tag",
	},
	{
		Name:        "image_tags",
		Type:        "string",
		Description: "Override the default image tags (json-encoded map)",
		Default:     cty.StringVal("{}"),
	},
	{
		Name:        "stack_name",
		Type:        "string",
		Description: "Happy Path stack name",
	},
	{
		Name:        "wait_for_steady_state",
		Type:        "bool",
		Description: "Should terraform block until k8s deployment reaches a steady state?",
		Default:     cty.BoolVal(true),
	},
}

const requiredTerraformVersion = ">= 1.3"

type TfGenerator struct {
	happyConfig *config.HappyConfig
}

func NewTfGenerator(happyConfig *config.HappyConfig) TfGenerator {
	return TfGenerator{
		happyConfig: happyConfig,
	}
}

func (tf *TfGenerator) GenerateMain(srcDir, moduleSource string, vars map[string]*tfconfig.Variable) error {
	variables := []Variable{}
	for _, variable := range vars {
		expr, diags := hclsyntax.ParseExpression([]byte(variable.Type), "", hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			log.Errorf("Variable %s type cannot be parsed: %s", variable.Name, diags.Errs()[0].Error())
			continue
		}

		typ, typDefaults, diags := typeexpr.TypeConstraintWithDefaults(expr)
		if diags.HasErrors() {
			log.Errorf("Variable %s type cannot be evaluated: %s", variable.Name, diags.Errs()[0].Error())
			continue
		}

		defaultValue, err := gocty.ToCtyValue(variable.Default, typ)
		if err != nil {
			log.Errorf("Variable %s default value cannot be converted: %s", variable.Name, err.Error())
			continue
		}

		variables = append(variables, Variable{
			Name:           variable.Name,
			Description:    variable.Description,
			Type:           typ,
			ConstraintType: typ,
			Default:        defaultValue,
			TypeDefaults:   typDefaults,
		})
	}

	stackConfig, err := tf.happyConfig.GetStackConfig()
	if err != nil {
		return errors.Wrap(err, "Unable to get stack config")
	}

	tfFile, err := os.Create(filepath.Join(srcDir, "main.tf"))
	if err != nil {
		return errors.Wrap(err, "Unable to generate HCL code")
	}
	defer tfFile.Close()
	hclFile := hclwrite.NewEmptyFile()

	rootBody := hclFile.Body()
	rootBody.AppendUnstructuredTokens(comment(autogeneratedComment))
	moduleBlockBody := rootBody.AppendNewBlock("module", []string{"stack"}).Body()

	moduleBlockBody.SetAttributeValue("source", cty.StringVal(moduleSource))

	// Sort module variables alphabetically
	sort.SliceStable(variables, func(i, j int) bool {
		return strings.Compare(variables[i].Name, variables[j].Name) < 0
	})

	varMap := map[string]Variable{}

	for _, variable := range variables {
		varMap[variable.Name] = variable
	}

	// These module variables reference stack variables set by happy
	for _, v := range []string{"image_tag", "stack_name", "k8s_namespace"} {
		if _, ok := varMap[v]; !ok {
			continue
		}
		moduleBlockBody.SetAttributeRaw(v, tokens(fmt.Sprintf("var.%s", v)))
		delete(varMap, v)
	}

	// These module variables refer to stack variables
	if _, ok := varMap["image_tags"]; ok {
		moduleBlockBody.SetAttributeRaw("image_tags", tokens("jsondecode(var.image_tags)"))
		delete(varMap, "image_tags")
	}

	if _, ok := varMap["stack_prefix"]; ok {
		moduleBlockBody.SetAttributeRaw("stack_prefix", tokens("\"/${var.stack_name}\""))
		delete(varMap, "stack_prefix")
	}

	// These variable depend on the happy config
	if _, ok := varMap["app_name"]; ok {
		moduleBlockBody.SetAttributeValue("app_name", cty.StringVal(tf.happyConfig.App()))
		delete(varMap, "app_name")
	}

	if _, ok := varMap["deployment_stage"]; ok {
		moduleBlockBody.SetAttributeValue("deployment_stage", cty.StringVal(tf.happyConfig.GetEnv()))
		delete(varMap, "deployment_stage")
	}

	if variable, ok := varMap["services"]; ok {
		if !variable.Type.IsMapType() {
			return errors.Errorf("services variable must be an object type")
		}

		var serviceConfigs map[string]interface{}
		if sc, ok := stackConfig["services"]; ok {
			serviceConfigs = sc.(map[string]interface{})
		}

		values := map[string]cty.Value{}

		for _, service := range tf.happyConfig.GetServices() {
			var serviceConfig map[string]interface{}
			if sc, ok := serviceConfigs[service]; ok {
				serviceConfig = sc.(map[string]interface{})
			}

			values[service] = tf.generateServiceValues(variable, serviceConfig)
		}

		val := cty.MapVal(values)
		moduleBlockBody.SetAttributeValue(variable.Name, val)
		delete(varMap, "services")
	}

	// All other variables
	for _, variable := range variables {
		if _, ok := varMap[variable.Name]; !ok {
			continue
		}

		var value cty.Value
		if configuredValue, ok := stackConfig[variable.Name]; ok {
			if configuredValue != nil {
				var err error
				value, err = gocty.ToCtyValue(configuredValue, variable.ConstraintType)
				if err != nil {
					log.Infof("Unable to convert a parameter value (%s): %s; will use default.", variable.Name, err.Error())
				}
			}
		}

		if value.IsNull() && !variable.Default.IsNull() {
			value = variable.Default
		}

		moduleBlockBody.SetAttributeValue(variable.Name, value)
	}

	_, err = tfFile.Write(hclFile.Bytes())

	return err
}

func (tf *TfGenerator) generateServiceValues(variable Variable, serviceConfig map[string]interface{}) cty.Value {
	defaultValues := variable.TypeDefaults.Children[""].DefaultValues
	elem := map[string]cty.Value{}

	// Sort the service attributes alphabetically
	attributeNames := reflect.ValueOf(variable.Type.ElementType().AttributeTypes()).MapKeys()
	sort.SliceStable(attributeNames, func(i, j int) bool {
		return strings.Compare(attributeNames[i].String(), attributeNames[j].String()) < 0
	})

	for i := range attributeNames {
		k := attributeNames[i].String()
		if _, ok := elem[k]; !ok {
			var value cty.Value
			var defaultValue cty.Value

			// Look up service attributes in happy config
			if configuredValue, ok := serviceConfig[k]; ok {
				if configuredValue != nil {
					var err error
					value, err = gocty.ToCtyValue(configuredValue, variable.Type.ElementType().AttributeTypes()[k])
					if err != nil {
						log.Errorf("Unable to convert a parameter value (%s): %s; will use default.", k, err.Error())
					}
				}
			}

			// If nothing is configured in happy config, use module defaults
			if value.IsNull() {
				if def, ok := defaultValues[k]; ok {
					defaultValue = def
				}
			}

			if value.IsNull() && !defaultValue.IsNull() {
				value = defaultValue
			}

			elem[k] = value
		}
	}

	return cty.ObjectVal(elem)
}

func (tf *TfGenerator) GenerateProviders(srcDir string) error {
	tfFile, err := os.Create(filepath.Join(srcDir, "providers.tf"))
	if err != nil {
		return errors.Wrap(err, "Unable to generate HCL code")
	}
	defer tfFile.Close()
	hclFile := hclwrite.NewEmptyFile()

	rootBody := hclFile.Body()
	rootBody.AppendUnstructuredTokens(comment(autogeneratedComment))
	err = tf.generateAwsProvider(rootBody, "", "${var.aws_account_id}", "${var.aws_role}")
	if err != nil {
		return errors.Wrap(err, "Unable to generate HCL code for AWS provider")
	}

	err = tf.generateAwsProvider(rootBody, "czi-si", "626314663667", "tfe-si")
	if err != nil {
		return errors.Wrap(err, "Unable to generate HCL code for AWS provider")
	}

	eksBody := rootBody.AppendNewBlock("data", []string{"aws_eks_cluster", "cluster"}).Body()
	eksBody.SetAttributeRaw("name", tokens("var.k8s_cluster_id"))
	eksAuthBody := rootBody.AppendNewBlock("data", []string{"aws_eks_cluster_auth", "cluster"}).Body()
	eksAuthBody.SetAttributeRaw("name", tokens("var.k8s_cluster_id"))

	kubernetesProviderBody := rootBody.AppendNewBlock("provider", []string{"kubernetes"}).Body()
	kubernetesProviderBody.SetAttributeRaw("host", tokens("data.aws_eks_cluster.cluster.endpoint"))
	kubernetesProviderBody.SetAttributeRaw("cluster_ca_certificate", tokens("base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)"))
	kubernetesProviderBody.SetAttributeRaw("token", tokens("data.aws_eks_cluster_auth.cluster.token"))

	kubeNamespaceBody := rootBody.AppendNewBlock("data", []string{"kubernetes_namespace", "happy-namespace"}).Body()
	kubeNamespaceBody.AppendNewBlock("metadata", nil).Body().SetAttributeRaw("name", tokens("var.k8s_namespace"))

	awsAliasTokens := tokens("aws.czi-si")
	appKeyBody := rootBody.AppendNewBlock("data", []string{"aws_ssm_parameter", "dd_app_key"}).Body()
	appKeyBody.SetAttributeValue("name", cty.StringVal("/shared-infra-prod-datadog/app_key"))
	appKeyBody.SetAttributeRaw("provider", awsAliasTokens)
	apiKeyBody := rootBody.AppendNewBlock("data", []string{"aws_ssm_parameter", "dd_api_key"}).Body()
	apiKeyBody.SetAttributeValue("name", cty.StringVal("/shared-infra-prod-datadog/api_key"))
	apiKeyBody.SetAttributeRaw("provider", awsAliasTokens)

	datadogProviderBody := rootBody.AppendNewBlock("provider", []string{"datadog"}).Body()
	datadogProviderBody.SetAttributeRaw("app_key", tokens("data.aws_ssm_parameter.dd_app_key.value"))
	datadogProviderBody.SetAttributeRaw("api_key", tokens("data.aws_ssm_parameter.dd_api_key.value"))

	_, err = tfFile.Write(hclFile.Bytes())

	return err
}

func (tf TfGenerator) generateAwsProvider(rootBody *hclwrite.Body, alias, accountIdExpr, roleExpr string) error {
	awsProviderBody := rootBody.AppendNewBlock("provider", []string{"aws"}).Body()
	if alias != "" {
		awsProviderBody.SetAttributeValue("alias", cty.StringVal(alias))
	}
	awsProviderBody.SetAttributeValue("region", cty.StringVal(*tf.happyConfig.AwsRegion()))

	assumeRoleBlockBody := awsProviderBody.AppendNewBlock("assume_role", nil).Body()
	assumeRoleBlockBody.SetAttributeRaw("role_arn", tokens(fmt.Sprintf("\"arn:aws:iam::%s:role/%s\"", accountIdExpr, roleExpr)))
	awsProviderBody.SetAttributeRaw("allowed_account_ids", tokens(fmt.Sprintf("[\"%s\"]", accountIdExpr)))
	return nil
}

func (tf *TfGenerator) GenerateVersions(srcDir string) error {
	tfFile, err := os.Create(filepath.Join(srcDir, "versions.tf"))
	if err != nil {
		return errors.Wrap(err, "Unable to generate HCL code")
	}
	defer tfFile.Close()
	hclFile := hclwrite.NewEmptyFile()

	rootBody := hclFile.Body()
	rootBody.AppendUnstructuredTokens(comment(autogeneratedComment))
	terraformBlockBody := rootBody.AppendNewBlock("terraform", nil).Body()
	terraformBlockBody.SetAttributeValue("required_version", cty.StringVal(requiredTerraformVersion))
	requiredProvidersBody := terraformBlockBody.AppendNewBlock("required_providers", nil).Body()

	for _, provider := range requiredProviders {
		p := cty.ObjectVal(map[string]cty.Value{
			"source":  cty.StringVal(provider.Source),
			"version": cty.StringVal(provider.Version),
		})
		requiredProvidersBody.SetAttributeValue(provider.Name, p)

	}

	_, err = tfFile.Write(hclFile.Bytes())

	return err
}

func (tf *TfGenerator) GenerateOutputs(srcDir string, outs map[string]*tfconfig.Output) error {
	outputs := []tfconfig.Output{}
	for _, output := range outs {
		outputs = append(outputs, *output)
	}

	tfFile, err := os.Create(filepath.Join(srcDir, "outputs.tf"))
	if err != nil {
		return errors.Wrap(err, "Unable to generate HCL code")
	}
	defer tfFile.Close()
	hclFile := hclwrite.NewEmptyFile()

	rootBody := hclFile.Body()
	rootBody.AppendUnstructuredTokens(comment(autogeneratedComment))

	sort.SliceStable(outputs, func(i, j int) bool {
		return strings.Compare(outputs[i].Name, outputs[j].Name) < 0
	})

	for _, output := range outputs {
		moduleOutputBody := rootBody.AppendNewBlock("output", []string{output.Name}).Body()
		if len(output.Description) > 0 {
			moduleOutputBody.SetAttributeValue("description", cty.StringVal(output.Description))
		}
		moduleOutputBody.SetAttributeValue("sensitive", cty.BoolVal(output.Sensitive))
		tokens := hclwrite.TokensForTraversal(hcl.Traversal{
			hcl.TraverseRoot{Name: "module"},
			hcl.TraverseAttr{Name: "stack"},
			hcl.TraverseAttr{Name: output.Name},
		})
		moduleOutputBody.SetAttributeRaw("value", tokens)
	}

	_, err = tfFile.Write(hclFile.Bytes())

	return err
}

func (tf *TfGenerator) GenerateVariables(srcDir string) error {
	tfFile, err := os.Create(filepath.Join(srcDir, "variables.tf"))
	if err != nil {
		return errors.Wrap(err, "Unable to generate HCL code")
	}
	defer tfFile.Close()
	hclFile := hclwrite.NewEmptyFile()

	rootBody := hclFile.Body()
	rootBody.AppendUnstructuredTokens(comment(autogeneratedComment))
	for _, variable := range requiredVariables {
		variableBody := rootBody.AppendNewBlock("variable", []string{variable.Name}).Body()
		tokens := hclwrite.TokensForTraversal(hcl.Traversal{
			hcl.TraverseRoot{Name: variable.Type},
		})
		variableBody.SetAttributeRaw("type", tokens)
		variableBody.SetAttributeValue("description", cty.StringVal(variable.Description))
		if !variable.Default.IsNull() {
			variableBody.SetAttributeValue("default", variable.Default)
		}
	}

	_, err = tfFile.Write(hclFile.Bytes())

	return err
}

func ParseModuleSource(moduleSource string) (gitUrl string, modulePath string, ref string, err error) {

	parts := strings.Split(moduleSource, "//")
	if len(parts) < 2 {
		return "", "", "", errors.Errorf("invalid module source %s", moduleSource)
	}

	gitUrl = parts[0]
	modulePathAndRef := parts[1]

	modulePathAndRefParts := strings.Split(modulePathAndRef, "?ref=")

	if len(modulePathAndRefParts) < 2 {
		return "", "", "", errors.Errorf("invalid module source, reference is missing %s", moduleSource)
	}

	modulePath = modulePathAndRefParts[0]
	ref = modulePathAndRefParts[1]

	return gitUrl, modulePath, ref, nil
}

func stringToTokens(value string) (hclwrite.Tokens, error) {
	file, diags := hclwrite.ParseConfig([]byte("attr = "+value), "", hcl.Pos{Line: 1, Column: 1})
	if diags.HasErrors() {
		return nil, diags.Errs()[0]
	}
	attr := file.Body().GetAttribute("attr")
	return attr.Expr().BuildTokens(hclwrite.Tokens{}), nil
}

func tokens(value string) hclwrite.Tokens {
	tokens, err := stringToTokens(value)
	if err != nil {
		log.Errorf("Unable to parse an HCL expression: %s: %s", value, err.Error())
	}
	return tokens
}

func comment(value string) hclwrite.Tokens {
	return hclwrite.Tokens{
		&hclwrite.Token{
			Type:         hclsyntax.TokenComment,
			Bytes:        []byte(value),
			SpacesBefore: 0,
		},
		&hclwrite.Token{
			Type:         hclsyntax.TokenNewline,
			Bytes:        []byte("\n"),
			SpacesBefore: 0,
		},
	}
}
