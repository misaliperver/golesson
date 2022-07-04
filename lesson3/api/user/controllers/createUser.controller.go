package controllers

import "github.com/gofiber/fiber/v2"

func CreateUser(c *fiber.Ctx) error {
	return c.Status(200).SendString("Create user")
}
