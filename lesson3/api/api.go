package api

import (
	"github.com/gofiber/fiber/v2"
	questionApi "github.com/misaliperver/golesson/lesson3/api/question"
	userApi "github.com/misaliperver/golesson/lesson3/api/user"
)

func RouterInitialize(app *fiber.App) {
	apiGroup := app.Group("/api") // /api

	userApi.Initialize(apiGroup)
	questionApi.Initialize(apiGroup)

}
