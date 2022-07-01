package questionApi

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misaliperver/golesson/lesson3/lib/logger"
)

var (
	l = logger.NewLogger()
)

func Initialize(parrentGroup fiber.Router) {
	group := parrentGroup.Group("/question") // /api/question

	group.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Question Routing!")
	})

	l.Debug("QuestionRouter initialized: ", group, "from parent:", parrentGroup)
}
