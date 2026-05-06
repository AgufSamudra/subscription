package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/AgufSamudra/subscription/src/internal/apperror"
)

func parseAndValidateJWT(c *fiber.Ctx, secretEnvKey string) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return apperror.UnauthorizedError("Missing Authorization header", nil)
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		return apperror.UnauthorizedError("Invalid Authorization header format", nil)
	}
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	// Langsung baca dari env yang sudah di-load saat startup
	jwtSecret := os.Getenv(secretEnvKey)
	if jwtSecret == "" {
		return apperror.InternalServerError("JWT secret is not set", nil)
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperror.UnauthorizedError("Unexpected signing method", nil)
		}
		return []byte(jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return apperror.UnauthorizedError("Invalid or expired token", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return apperror.UnauthorizedError("Invalid token claims", nil)
	}

	c.Locals("user", claims)
	if userID, ok := claims["sub"].(string); ok {
		c.Locals("user_id", userID)
	}

	return c.Next()
}

func JWTMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return parseAndValidateJWT(c, "ACCESS_TOKEN_SECRET")
	}
}
