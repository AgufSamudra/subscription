package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/AgufSamudra/subscription/src/internal/apperror"
	"github.com/AgufSamudra/subscription/src/internal/databases"
	"github.com/AgufSamudra/subscription/src/internal/routes"
	"github.com/AgufSamudra/subscription/src/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func main() {
	if err := utils.LoadEnv(); err != nil {
		utils.Fatalf("gagal load file .env: %v", err)
	}

	dbClient, err := databases.PostgreSQLConnection(context.Background())
	if err != nil {
		utils.Fatalf("aplikasi dibatalkan start karena koneksi database gagal: %v", err)
	}

	defer func() {
		if err := dbClient.Close(); err != nil {
			utils.Errorf("gagal menutup koneksi database: %v", err)
		}
	}()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			var appErr *apperror.AppError
			if errors.As(err, &appErr) {
				return c.Status(appErr.StatusCode).JSON(fiber.Map{
					"error":   true,
					"code":    appErr.StatusCode,
					"message": appErr.Message,
				})
			}

			var fiberErr *fiber.Error
			if errors.As(err, &fiberErr) {
				return c.Status(fiberErr.Code).JSON(fiber.Map{
					"error":   true,
					"code":    fiberErr.Code,
					"message": fiberErr.Message,
				})
			}

			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error":   true,
				"code":    fiber.StatusInternalServerError,
				"message": "internal server error",
			})
		},
	})

	routes.RegisterRoutes(app, dbClient)

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = os.Getenv("PORT")
	}
	if port == "" {
		utils.Fatal("APP_PORT atau PORT wajib diisi")
	}

	fmt.Printf("Server running at http://localhost:%s\n", port)
	if err := app.Listen(":" + port); err != nil {
		utils.Fatalf("server gagal dijalankan: %v", err)
	}
}
