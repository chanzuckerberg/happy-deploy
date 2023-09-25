package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/chanzuckerberg/happy/api/pkg/setup"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	sqsMaxMessages     int32 = 10
	sqsPollWaitSeconds int32 = 10
)

func init() {
	rootCmd.AddCommand(eventConsumerCmd)
}

var eventConsumerCmd = &cobra.Command{
	Use:          "event-consumer",
	Short:        "run the happy api server",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		logrus.Info("Running event consumer")
		cfg := setup.GetConfiguration()
		cfg.LogConfiguration()

		go startHealthApp(cfg)

		awsCfg, err := getAwsConfig(
			cmd.Context(),
			cfg.EventConsumer.AWSProfile,
			getRegion(cfg.EventConsumer.QueueURL),
		)
		if err != nil {
			return errors.Wrap(err, "failed to get aws config")
		}
		sqsSvc := sqs.NewFromConfig(*awsCfg)

		chnMessages := make(chan *types.Message, sqsMaxMessages)
		go pollSqs(cmd.Context(), chnMessages, sqsSvc, cfg)

		logrus.Infof("Listening on queue: %s", cfg.EventConsumer.QueueURL)

		for message := range chnMessages {
			handleMessage(message, sqsSvc, cfg)
			deleteMessage(cmd.Context(), message, sqsSvc, cfg)
		}

		return nil
	},
}

func pollSqs(ctx context.Context, chn chan<- *types.Message, sqsSvc *sqs.Client, cfg *setup.Configuration) {
	for {
		output, err := sqsSvc.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl:            &cfg.EventConsumer.QueueURL,
			MaxNumberOfMessages: *aws.Int32(sqsMaxMessages),
			WaitTimeSeconds:     *aws.Int32(sqsPollWaitSeconds),
		})

		if err != nil {
			logrus.Errorf("Failed to fetch sqs messages: %v", err)
		}

		for _, message := range output.Messages {
			chn <- &message
		}
	}
}

func deleteMessage(ctx context.Context, message *types.Message, sqsSvc *sqs.Client, cfg *setup.Configuration) {
	_, err := sqsSvc.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(cfg.EventConsumer.QueueURL),
		ReceiptHandle: message.ReceiptHandle,
	})

	if err != nil {
		logrus.Errorf("Failed to delete sqs message %v", err)
	}
}

func handleMessage(message *types.Message, sqsSvc *sqs.Client, cfg *setup.Configuration) {
	logrus.Infof("Received message: %s", *message.Body)

	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(*message.Body), &result)
	if err != nil {
		logrus.Errorf("Failed to unmarshal message %v", err)
	}

	logrus.Infof("--> Parsed message: %s", result["Message"])

	// TODO: parse message to proto and save in DB
}

func getAwsConfig(ctx context.Context, profile, region string) (*aws.Config, error) {
	opts := []func(*config.LoadOptions) error{}
	if len(profile) > 0 {
		opts = append(opts, config.WithSharedConfigProfile(profile))
	}
	if len(region) > 0 {
		opts = append(opts, config.WithRegion(region))
	}

	cfg, err := config.LoadDefaultConfig(ctx, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load config")
	}

	return &cfg, nil
}

func getRegion(queueUrl string) string {
	res := strings.SplitN(queueUrl, ".", 3)
	region := res[1]
	if len(region) > 0 {
		return region
	}
	return "us-west-2"
}

// This exists to fulfil the healthcheck requirement of the load balancer
func startHealthApp(cfg *setup.Configuration) {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	// one greater than the main server port to avoid conflicts when running simultaneously
	app.Listen(fmt.Sprintf(":%d", cfg.Api.Port+1))
}
