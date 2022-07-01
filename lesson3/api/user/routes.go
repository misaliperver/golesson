package userApi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misaliperver/golesson/lesson3/lib/logger"
)

var (
	l = logger.NewLogger("useApi", "api/user/routes.go")
)

func Initialize(parrentGroup fiber.Router) {
	group := parrentGroup.Group("/user") // /api/user

	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("User Routing!")
	})

	l.Debug("UserRouter initialized: ", group, "from parent:", parrentGroup)
}
