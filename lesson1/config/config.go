package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type db struct {
	uri  string
	name string
}

type env struct {
	defaultLang   string
	appListenPort string
	db            db
	consoleLevel  string
}

func initialize() env {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return env{
		defaultLang:   os.Getenv("DEFAULT_LANG"),
		appListenPort: os.Getenv("APP_LISTEN_PORT"),
		db: db{
			uri:  os.Getenv("MONGO_URI"),
			name: os.Getenv("DBNAME"),
		},
		consoleLevel: os.Getenv("CONSOLE_LEVEL"),
	}
}
