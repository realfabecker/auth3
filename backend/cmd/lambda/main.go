package main

import (
	"context"
	"errors"

	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gofiber/fiber/v2"

	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/realfabecker/auth3/internal/adapters"
	cordom "github.com/realfabecker/auth3/internal/core/domain"
	corpts "github.com/realfabecker/auth3/internal/core/ports"
)

var fiberLambda *fiberadapter.FiberLambda

func init() {
	log.Println("lambda cold start")
	if err := adapters.Container.Invoke(func(
		app corpts.HttpHandler,
		walletConfig *cordom.Config,
	) error {
		log.Println("app:route-register")
		if err := app.Register(); err != nil {
			return err
		}

		log.Println("app:app-capture")
		fapp, b := app.GetApp().(*fiber.App)
		if !b {
			return errors.New("unable to get fiber instance")
		}

		log.Println("app:app-adapter")
		fiberLambda = fiberadapter.New(fapp)
		return nil
	}); err != nil {
		log.Fatalln(err)
	}
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Processing Pre Request Proxy")
	response, err := fiberLambda.ProxyWithContext(ctx, events.APIGatewayProxyRequest{
		Path:                  request.Path,
		Body:                  request.Body,
		Headers:               request.Headers,
		HTTPMethod:            request.HTTPMethod,
		QueryStringParameters: request.QueryStringParameters,
	})
	if response.Headers == nil {
		response.Headers = map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "*",
			"Access-Control-Allow-Headers": "*",
		}
	} else {
		response.Headers["Access-Control-Allow-Origin"] = "*"
		response.Headers["Access-Control-Allow-Methods"] = "*"
		response.Headers["Access-Control-Allow-Headers"] = "*"
	}
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: response.StatusCode,
		Headers:    response.Headers,
		Body:       response.Body,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
