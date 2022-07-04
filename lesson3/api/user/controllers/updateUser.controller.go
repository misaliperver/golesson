package controllers

import "github.com/gofiber/fiber/v2"

func UpdateUser(c *fiber.Ctx) error {
	return c.Status(200).SendString("Update user")
}
