package controllers

import "github.com/gofiber/fiber/v2"

func DeleteUser(c *fiber.Ctx) error {
	return c.Status(200).SendString("Delete user")
}
