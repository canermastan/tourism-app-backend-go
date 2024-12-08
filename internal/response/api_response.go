package response

import (
	"github.com/gofiber/fiber/v2"
)

type ApiResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, data interface{}, message string) error {
	return c.Status(200).JSON(ApiResponse{
		Status:  "true",
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(ApiResponse{
		Status:  "false",
		Message: message,
	})
}
