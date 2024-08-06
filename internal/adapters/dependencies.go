package adapters

import (
	"MiddlewareAuth/cmd/config"
	auth "MiddlewareAuth/internal/adapters/repository/auth"
	repository "MiddlewareAuth/internal/core/domain/repository"
	"MiddlewareAuth/internal/core/services"
	"fmt"
	"net/http"

	"MiddlewareAuth/cmd/logging"
)

type Dependencies struct {
	AuthService *services.AppAuthService
}

func InitDependencies(cfg *config.Config, logger logging.Logger, httpClient *http.Client) *Dependencies {
	redisClient, err := auth.NewRedisClient(cfg.Redis.RedisAddr, cfg.Redis.RedisPassword)
	if err != nil {
		logger.Error("Error al crear el cliente Redis")
		fmt.Println(err)
	} else {
		logger.Info("Cliente Redis creado exitosamente")
	}
	userRepository := repository.NewRedisUserRepository(redisClient)
	authService := services.NewAuthService(*userRepository)

	return &Dependencies{
		AuthService: authService,
	}
}
