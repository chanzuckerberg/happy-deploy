package tf

import (
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/ext/typeexpr"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/pkg/errors"
	"github.com/zclconf/go-cty/cty"
)

type TfParser struct {
}

func NewTfParser() TfParser {
	return TfParser{}
}

var moduleBlockSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "module",
			LabelNames: []string{"name"},
		},
	},
}

var variableBlockSchema = &hcl.BodySchema{
	Attributes: []hcl.AttributeSchema{
		{
			Name: "description",
		},
		{
			Name: "default",
		},
		{
			Name: "type",
		},
		{
			Name: "sensitive",
		},
		{
			Name: "nullable",
		},
	},
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type: "validation",
		},
	},
}

var outputBlockSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{
			Type:       "output",
			LabelNames: []string{"name"},
		},
	},
}

func (tf TfParser) ParseServices(dir string) (map[string]bool, error) {
	var services map[string]bool = map[string]bool{}

	err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			if d.Name() == ".terraform" || d.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}
		if filepath.Ext(path) != ".tf" {
			return nil
		}

		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		f, diags := hclsyntax.ParseConfig(b, path, hcl.Pos{Line: 1, Column: 1})
		if diags.HasErrors() {
			return errors.Wrapf(diags.Errs()[0], "failed to parse %s", path)
		}

		content, _, diags := f.Body.PartialContent(moduleBlockSchema)
		if diags.HasErrors() {
			return errors.New("Terraform code has errors")
		}

		for _, block := range content.Blocks {
			if block.Type != "module" {
				continue
			}

			attrs, diags := block.Body.JustAttributes()
			if diags.HasErrors() {
				return errors.New("Terraform code has errors")
			}
			var sourceAttr *hcl.Attribute
			var ok bool
			if sourceAttr, ok = attrs["source"]; !ok {
				// Module without a source
				continue
			}

			source, diags := sourceAttr.Expr.(*hclsyntax.TemplateExpr).Parts[0].Value(nil)
			if diags.HasErrors() {
				return errors.New("Terraform code has errors")
			}

			if !strings.Contains(source.AsString(), "modules/happy-stack-eks") && !strings.Contains(source.AsString(), "modules/happy-stack-ecs") {
				// Not a happy stack module
				continue
			}

			if servicesAttr, ok := attrs["services"]; ok {
				switch servicesAttr.Expr.(type) {
				case *hclsyntax.ObjectConsExpr:
					for _, item := range servicesAttr.Expr.(*hclsyntax.ObjectConsExpr).Items {
						key, _ := item.KeyExpr.Value(nil)
						services[key.AsString()] = true
					}
				}
			}
		}

		return nil
	})
	if err != nil {
		return services, errors.Wrap(err, "failed to parse terraform files")
	}
	return services, nil
}

type Value struct {
	Value  cty.Value
	Source string
	Expr   hcl.Expression
	Range  hcl.Range
}

type Variable struct {
	Name        string
	Description string
	Default     cty.Value

	Type           cty.Type
	ConstraintType cty.Type
	TypeDefaults   *typeexpr.Defaults
}
