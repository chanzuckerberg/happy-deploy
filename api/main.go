package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_ "github.com/chanzuckerberg/happy/api/docs" // import API docs
	"github.com/chanzuckerberg/happy/api/pkg/api"
	"github.com/chanzuckerberg/happy/api/pkg/dbutil"
	"github.com/chanzuckerberg/happy/api/pkg/setup"
	sentry "github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

func exec(ctx context.Context) error {
	cfg := setup.GetConfiguration()

	m, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	logrus.Info("Running with configuration:\n", string(m))

	err = sentry.Init(sentry.ClientOptions{
		Dsn:              cfg.Sentry.DSN,
		Environment:      cfg.Api.DeploymentStage,
		EnableTracing:    true,
		TracesSampleRate: 1.0,
	})
	if err == nil {
		logrus.Info("Sentry enabled for environment: ", cfg.Api.DeploymentStage)
		// Flush buffered events before the program terminates.
		// Set the timeout to the maximum duration the program can afford to wait.
		defer sentry.Flush(2 * time.Second)
	} else {
		logrus.Info("Sentry disabled for environment: ", cfg.Api.DeploymentStage)
	}

	// run the DB migrations
	dbutil.MakeDB(cfg.Database).AutoMigrate()

	// create a mux to route requests to the correct app
	rootMux := http.NewServeMux()

	// create the Fiber app
	app := api.MakeApp(ctx, cfg)
	nativeHandler := adaptor.FiberApp(app.FiberApp)
	rootMux.Handle("/v1/", http.StripPrefix("/v1", nativeHandler))

	// create the Ogent app
	// uncomment the following to enable and test ent but do not commit because we don't have auth plugged in yet
	// svr, err := api.GetOgentServer(cfg)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// rootMux.Handle("/v2/", http.StripPrefix("/v2", svr))

	return http.ListenAndServe(fmt.Sprintf(":%d", cfg.Api.Port), rootMux)
}

// @title       Happy API
// @description An API to encapsulate Happy Path functionality
// @BasePath    /
func main() {
	err := exec(context.Background())
	if err != nil {
		logrus.Error(err)
	}
}
