package controllers

import (
	helpers "github.com/misaliperver/golesson/lesson3/api/_helpers"
	"github.com/misaliperver/golesson/lesson3/lib/logger"
	"github.com/misaliperver/golesson/lesson3/models"

	"github.com/gofiber/fiber/v2"
)

var (
	l = logger.NewLogger()
)

func CreateUser(c *fiber.Ctx) error {
	newUser := new(models.User)

	if err := c.BodyParser(newUser); err != nil {
		l.Error("Error parsing body: ", err)

		return helpers.BadRequest(c, "Error parsing body", err)
	}

	err := models.User.Create(newUser)

	if err != nil {
		l.Error("Error creating user: ", err)

		return helpers.ServerError(c, "User created failed", err)
	}

	return helpers.Ok(c, "User created successfully", newUser)
}
