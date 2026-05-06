package handlers

import (
	"database/sql"

	"github.com/AgufSamudra/subscription/src/internal/apperror"
	"github.com/AgufSamudra/subscription/src/internal/models"
	"github.com/AgufSamudra/subscription/src/internal/services"
	"github.com/AgufSamudra/subscription/src/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	db *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{db: db}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var request models.Register
	if err := c.BodyParser(&request); err != nil {
		return apperror.BadRequestError("invalid request body", err)
	}

	if err := utils.ValidateStruct(request); err != nil {
		return err
	}

	// init service
	service, err := services.NewAuthService(h.db)
	if err != nil {
		return apperror.InternalError(err)
	}

	// call service
	result, err := service.RegisterService(c.Context(), request.Email, request.Password)
	if err != nil {
		return apperror.InternalError(err)
	}

	return utils.Success(c, "Register Successfully!", result)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var request models.Login
	if err := c.BodyParser(&request); err != nil {
		return apperror.BadRequestError("invalid request body", err)
	}

	if err := utils.ValidateStruct(request); err != nil {
		return err
	}

	return nil
}
