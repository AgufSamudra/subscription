package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const accessTokenSecretEnvKey = "ACCESS_TOKEN_SECRET"

func GenerateJWT(subject string, expiresIn time.Duration, additionalClaims map[string]interface{}) (string, error) {
	return GenerateJWTWithSecretEnv(accessTokenSecretEnvKey, subject, expiresIn, additionalClaims)
}

func GenerateJWTWithSecretEnv(secretEnvKey, subject string, expiresIn time.Duration, additionalClaims map[string]interface{}) (string, error) {
	secret := os.Getenv(secretEnvKey)
	if secret == "" {
		return "", errors.New("jwt secret is not set")
	}

	if expiresIn <= 0 {
		return "", errors.New("jwt expiration duration must be greater than zero")
	}

	now := time.Now()
	claims := jwt.MapClaims{
		"sub": subject,
		"iat": now.Unix(),
		"exp": now.Add(expiresIn).Unix(),
	}

	for key, value := range additionalClaims {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
