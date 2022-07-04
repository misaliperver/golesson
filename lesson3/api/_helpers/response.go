package helpers

import "github.com/gofiber/fiber/v2"

func Ok(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
		"success": true,
		"message": message,
		"data":   data,
	})
}


func BadRequest(c *fiber.Ctx, message string, err error) error {
	return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
		"success": false,
		"message": message,
		"error":   err,
	})
}

func ServerError(c *fiber.Ctx, message string, err error) error {
	return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
		"success": false,
		"message": message
		"error":   err,
	})
}
