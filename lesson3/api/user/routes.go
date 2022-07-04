package userApi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misaliperver/golesson/lesson3/api/user/controllers"
	"github.com/misaliperver/golesson/lesson3/lib/logger"
)

var (
	l = logger.NewLogger()
)

func Initialize(parrentGroup fiber.Router) {
	group := parrentGroup.Group("/user") // /api/user

	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("User Routing!")
	})

	group.Get("/list", controllers.ListUser)
	group.Post("/create", controllers.CreateUser)
	group.Get("/:username", controllers.FindUser)
	group.Put("/:username", controllers.UpdateUser)
	group.Delete("/:username", controllers.DeleteUser)

	l.Debug("UserRouter initialized: ", group, "from parent:", parrentGroup)
}
