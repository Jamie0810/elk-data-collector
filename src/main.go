package main

import (
	"data-collector/config"
	"data-collector/handler"
	internal "data-collector/pkg/log"
	"log"

	"github.com/pkg/errors"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	config, err := config.InitConfig("./config")
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to initial config."))
	}

	logger, err := internal.InitLogger(config)
	if err != nil {
		log.Fatal(errors.Wrap(err, "Failed to initial logger."))
	}

	c := handler.InitLambdaController(config, logger)
	lambda.Start(c.LambdaHandler)
}
