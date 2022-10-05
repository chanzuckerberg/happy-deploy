package main

import (
	_ "github.com/chanzuckerberg/happy-api/docs" // import API docs
	"github.com/chanzuckerberg/happy-api/pkg/api"
	"github.com/chanzuckerberg/happy-api/pkg/setup"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func exec() error {
	config, err := setup.GetConfiguration()
	if err != nil {
		return err
	}

	m, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	logrus.Info("Running with configuration:\n", string(m))

	app := api.MakeApp(config)

	return app.Listen()
}

// @title       Happy API testing doc update...
// @description An API to encapsulate Happy Path functionality
// @BasePath    /
func main() {
	err := exec()
	if err != nil {
		logrus.Error(err)
	}
}
