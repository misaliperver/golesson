package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/misaliperver/golesson/lesson3/config"
)

var (
	env, _    = config.Get()
	ApiLogger = logger.New(logger.Config{
		Format:     env.LoggingFormat,
		TimeFormat: "02-Jan-2006T15:04:05",
	})
)

func NotFound(c *fiber.Ctx) error {
	return c.Status(404).SendString("Not found")
}
