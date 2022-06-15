package config

import (
	"errors"
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

var environment env

func setNewEnvironment() {
	environment = env{
		defaultLang:   os.Getenv("DEFAULT_LANG"),
		appListenPort: os.Getenv("APP_LISTEN_PORT"),
		db: db{
			uri:  os.Getenv("MONGO_URI"),
			name: os.Getenv("DBNAME"),
		},
		consoleLevel: os.Getenv("CONSOLE_LEVEL"),
	}
}

func Initialize() (env, error) {
	if environment != (env{}) {
		log.Fatal("[config] already initialized")
		return env{}, errors.New("[config] intialized already")
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("[config] error when loading .env file")
		return env{}, errors.New("[config] error when loading .env file")
	}

	log.Print("[config] Initialized .env file")

	setNewEnvironment()

	return environment, nil
}

func Get() (env, error) {
	if environment == (env{}) {
		return env{}, errors.New("[config] not initialized")
	}

	return environment, nil
}
