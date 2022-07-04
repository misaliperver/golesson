package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/misaliperver/golesson/lesson3/api"
	middleware "github.com/misaliperver/golesson/lesson3/api/_middleware"
	"github.com/misaliperver/golesson/lesson3/config"
	"github.com/misaliperver/golesson/lesson3/lib/logger"
	"github.com/misaliperver/golesson/lesson3/models/user"
)

var (
	l = logger.NewLogger()
)

func main() {
	env, _ := config.Get()
	finded, _ := user.FindAll()

	l.Info("Finded collection count:", len(finded))

	// create app
	app := fiber.New(fiber.Config{
		Prefork:       env.AppPrefork,
		CaseSensitive: false,
		StrictRouting: true,
		ServerHeader:  "Lesson3",
		AppName:       "GoLesson App v0.0.1",
	})

	// app init pre handlers (middlewares)
	app.Use(middleware.ApiLogger)

	// app init api routes
	api.RouterInitialize(app)

	// app init error handlers (middlewares)
	app.Use(middleware.NotFound)

	// app listen on env.ApplistenPort
	app.Listen(env.AppListenPort)
}
