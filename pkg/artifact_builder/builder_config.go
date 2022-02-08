package artifact_builder

import (
	"github.com/chanzuckerberg/happy/pkg/config"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

type ServiceBuild struct {
	Dockerfile string `yaml:"dockerfile"`
}

type ServiceConfig struct {
	Image   string                 `yaml:"image"`
	Build   *ServiceBuild          `yaml:"build"`
	Network map[string]interface{} `yaml:"networks"`
}

type ConfigData struct {
	Services map[string]ServiceConfig `yaml:"services"`
}

type BuilderConfig struct {
	composeFile string
	envFile     string
	dockerRepo  string

	// parse the passed in config file and populate some fields
	configData *ConfigData
}

func NewBuilderConfig(bootstrap *config.Bootstrap, happyConfig *config.HappyConfig) *BuilderConfig {
	return &BuilderConfig{
		composeFile: bootstrap.DockerComposeConfigPath,
		envFile:     happyConfig.GetComposeEnvFile(),
		dockerRepo:  happyConfig.GetDockerRepo(),
	}
}

func (s *BuilderConfig) GetContainers() []string {
	var containers []string
	configData, _ := s.getConfigData()
	for _, service := range configData.Services {
		for _, network := range service.Network {
			for _, aliases := range network.(map[interface{}]interface{}) {
				for _, alias := range aliases.([]interface{}) {
					containers = append(containers, alias.(string))
				}
			}
		}
	}

	return containers
}

func (s *BuilderConfig) getConfigData() (*ConfigData, error) {
	if s.configData != nil {
		return s.configData, nil
	}

	configFile, err := InvokeDockerCompose(*s, "config")
	if err != nil {
		return nil, errors.Wrap(err, "process failure")
	}

	var configData ConfigData
	err = yaml.Unmarshal(configFile, &configData)
	if err != nil {
		return nil, errors.Wrap(err, "fail to parse yaml")
	}
	s.configData = &configData

	return s.configData, nil
}

func (s *BuilderConfig) GetBuildEnv() []string {
	dockerRepoStr := "DOCKER_REPO=" + s.dockerRepo

	return []string{
		"DOCKER_BUILDKIT=1",
		"BUILDKIT_INLINE_CACHE=1",
		"COMPOSE_DOCKER_CLI_BUILD=1",
		dockerRepoStr,
	}
}

func (s *BuilderConfig) GetBuildServicesImage() (map[string]string, error) {
	configData, err := s.getConfigData()
	if err != nil {
		return nil, err
	}

	svcs := map[string]string{}
	for serviceName, service := range configData.Services {
		if service.Build != nil {
			svcs[serviceName] = service.Image
		}
	}
	return svcs, nil
}
