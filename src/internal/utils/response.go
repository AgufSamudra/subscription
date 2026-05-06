package utils

import "github.com/gofiber/fiber/v2"

type SuccessResponse struct {
	Error   bool        `json:"error"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Error   bool   `json:"error"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Success(c *fiber.Ctx, message string, data interface{}) error {
	statusCode := fiber.StatusOK

	return c.Status(statusCode).JSON(SuccessResponse{
		Error:   false,
		Code:    statusCode,
		Message: message,
		Data:    data,
	})
}
