package repository_models

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/redis/go-redis/v9"
)

type RedisUserRepository struct {
	RedisClient *redis.Client
}

func NewRedisUserRepository(client *redis.Client) *RedisUserRepository {
	// This is the constructor you were looking for
	return &RedisUserRepository{
		RedisClient: client,
	}
}

func (r *RedisUserRepository) Login(ctx context.Context, username, password string) (string, error) {

	// Generar el token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte("s3cr3tK3yF0rJWT!"))
	if err != nil {
		return "", err
	}

	// Guardar el token en Redis
	err = r.RedisClient.Set(ctx, tokenString, "valid", time.Minute*30).Err()
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
