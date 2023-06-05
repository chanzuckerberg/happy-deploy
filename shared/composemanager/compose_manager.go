package composemanager

import (
	"context"
	"encoding/json"
	"os"

	"github.com/chanzuckerberg/happy/shared/config"
	"github.com/compose-spec/compose-go/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type ComposeManager struct {
	HappyConfig *config.HappyConfig
}

func NewComposeManager() ComposeManager {
	return ComposeManager{}
}

func (c ComposeManager) WithHappyConfig(happyConfig *config.HappyConfig) ComposeManager {
	c.HappyConfig = happyConfig
	return c
}

func (c ComposeManager) Generate(ctx context.Context) error {
	p := types.Project{}
	p.Services = types.Services{}

	stackDef, err := c.HappyConfig.GetStackConfig()
	if err != nil {
		return errors.Wrap(err, "unable to get stack config")
	}

	_, ok := stackDef["services"]
	if !ok {
		return errors.New("unable to find services in stack config")
	}

	servicesDef := stackDef["services"].(map[string]any)
	if len(servicesDef) == 0 {
		return errors.New("no service settings are defined in stack config")
	}

	for _, service := range c.HappyConfig.GetData().Services {
		if sd, ok := servicesDef[service]; ok {
			serviceDef := sd.(map[string]any)
			serviceConfig := types.ServiceConfig{
				Name:     service,
				Image:    service,
				Profiles: []string{"*"},
				Build:    &types.BuildConfig{},
			}
			platform := serviceDef["platform_architecture"].(string)
			if len(platform) == 0 {
				platform = "linux/amd64"
			}

			jsonData, err := json.Marshal(serviceDef["build"])
			if err != nil {
				return errors.Wrap(err, "unable to marshal build config")
			}

			err = json.Unmarshal(jsonData, serviceConfig.Build)
			if err != nil {
				return errors.Wrap(err, "unable to unmarshal build config")
			}

			serviceConfig.Platform = platform
			p.Services = append(p.Services, serviceConfig)
		}
	}

	composeFilePath := c.HappyConfig.GetBootstrap().DockerComposeConfigPath
	logrus.Printf("Generating docker-compose.yml at %s", composeFilePath)
	configYaml, err := yaml.Marshal(p)
	if err != nil {
		return errors.Wrap(err, "unable to marshal compose config")
	}

	err = os.WriteFile(composeFilePath, configYaml, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "unable to write out %s", composeFilePath)
	}
	return nil
}
