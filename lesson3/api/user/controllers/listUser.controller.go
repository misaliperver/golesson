package controllers

import "github.com/gofiber/fiber/v2"

func ListUser(c *fiber.Ctx) error {
	return c.Status(200).SendString("List user")
}
