package main

import (
	"MiddlewareAuth/cmd/config"
	"MiddlewareAuth/internal/adapters"
	"MiddlewareAuth/internal/adapters/handlers"
	"context"
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	config2 "MiddlewareAuth/cmd/config"
	"MiddlewareAuth/cmd/config/model"
	"MiddlewareAuth/cmd/logging"
)

var (
	cfg        *config.Config
	httpClient *http.Client
	flagConfig = flag.String("config", "./cmd/configs/properties.yml", "path to the config file")

	logger                logging.Logger
	artifactResources     model.ArtifactResources
	flagArtifactResources = flag.String("flagArtifactResources", "./cmd/configs/resources.json", "path to the Resources file")

	levelLogging = os.Getenv("LEVEL_LOGGING")
	err          = errors.New("")
)

func init() {
	flag.Parse()

	artifactResources = config2.GetArtifactResources(*flagArtifactResources)

	logger = logging.New(levelLogging).With(context.TODO())

	// load application configurations
	cfg, err = config.Load(*flagConfig, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("failed to load application configuration %s", err))
		os.Exit(-1)
	}
}

func main() {
	tp := config.InitTracerProvider()
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Error(fmt.Sprintf("Error shutting down tracer provider: %v", err))
		}
	}()

	dependencies := adapters.InitDependencies(cfg, logger, httpClient)
	go handlers.CreateNewHttpServer(cfg, logger, artifactResources, dependencies)

	select {}
}
