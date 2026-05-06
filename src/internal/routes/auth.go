package routes

import (
	"database/sql"

	"github.com/AgufSamudra/subscription/src/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func authRoutes(app fiber.Router, db *sql.DB) {
	authRoutes := app.Group("/auth")

	init_handler := handlers.NewAuthHandler(db)

	// route
	authRoutes.Post("/login", init_handler.Login)
	authRoutes.Post("/register", init_handler.Register)
}
