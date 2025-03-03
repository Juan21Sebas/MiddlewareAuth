package handlers

import (
	"fmt"
	"os"

	"MiddlewareAuth/cmd/config"
	core_middleware "MiddlewareAuth/cmd/config/middleware"
	"MiddlewareAuth/cmd/config/model"
	"MiddlewareAuth/cmd/logging"
	"MiddlewareAuth/internal/adapters"
	middlewares "MiddlewareAuth/internal/adapters/handlers/http/middleware"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

var ServiceName = os.Getenv("OTEL_SERVICE_NAME")

func CreateNewHttpServer(cfg *config.Config, logger logging.Logger, artifactResources model.ArtifactResources, dependencies *adapters.Dependencies) {
	router := gin.New()
	router.Use(
		otelgin.Middleware(ServiceName),
		core_middleware.Logger(logger, &artifactResources, cfg.EnablePayloadLogging),
		middlewares.ErrorHandler(),
		middlewares.HandlerPanic(),
		core_middleware.SanitizeRequest(),
	)
	var (
		healthHttp = MakeNewHealthController()
		loansHttp  = MakeNewAuthHttp(*dependencies.AuthService)
	)
	app := router.Group(cfg.ContextPath)
	{
		HealthRoute(app, &healthHttp)
		LoansRouter(app, &loansHttp)
	}
	_ = router.Run(fmt.Sprintf(":%v", cfg.HttpServerPort))
}

func HealthRoute(router *gin.RouterGroup, endpoint *HealthHttp) {
	router.GET("/healthCheck", endpoint.HealthCheck())
}

func LoansRouter(router *gin.RouterGroup, endpoint *AuthHttp) {
	router.POST("/login", endpoint.Login())
}
