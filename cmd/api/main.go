package main

import (
	"fmt"
	"log"

	"bootstrap-echo-openapi-docker/api"
	"bootstrap-echo-openapi-docker/internal/server/v1"

	oapimiddleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

)

const port = 8112

func main() {
	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err, "error loading swagger spec")
	}

	// Create a server that satisfies the V1 API interface
	serverV1 := server.New()

	swagger.Servers = nil

	// Setup basic Echo router
	e := echo.New()

	// Stop Echo logging a useless, unstructured banner in the logs
	e.HideBanner = true

	// Generic middlewares we care about
	e.Use(middleware.CORS())

	// Use our validation middleware to check all requests against the OpenAPI schema.
	e.Use(oapimiddleware.OapiRequestValidator(swagger))

	// We now register our server above as the handler for the interface
	api.RegisterHandlers(e, serverV1)

	log.Printf("starting service on port %d", port)

	// And we serve HTTP until the world ends.
	log.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", port)), "failed to start service")
}

