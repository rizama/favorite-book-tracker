package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rizama/favorite-book-tracker/shared/utils"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Logic
	req_api_key := c.Get("x-api-key")
	api_key := os.Getenv("API_KEY")

	if req_api_key != api_key {
		resp, statuCode := utils.ResponseError(fiber.StatusForbidden, "Invalid API Key")
		return c.Status(statuCode).JSON(resp)
	}

	return c.Next()
}
