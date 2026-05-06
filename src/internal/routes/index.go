package routes

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, db *sql.DB) {

	api := app.Group("/api/v1")

	authRoutes(api, db)
}
