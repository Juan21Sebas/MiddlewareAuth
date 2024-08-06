package handlers

import "MiddlewareAuth/internal/core/services"

type AuthHttp struct {
	service services.AppAuthService
}

func MakeNewAuthHttp(service services.AppAuthService) AuthHttp {
	return AuthHttp{
		service: service,
	}
}
