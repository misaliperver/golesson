package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Db struct {
	uri  string
	name string
}

type Env struct {
	defaultLang   string
	appListenPort string
	db            Db
	consoleLevel  string
}

var environment Env

func setNewEnvironment() {
	environment = Env{
		defaultLang:   os.Getenv("DEFAULT_LANG"),
		appListenPort: os.Getenv("APP_LISTEN_PORT"),
		db: Db{
			uri:  os.Getenv("MONGO_URI"),
			name: os.Getenv("DBNAME"),
		},
		consoleLevel: os.Getenv("CONSOLE_LEVEL"),
	}
}

func Initialize() (*Env, error) {
	if environment != (Env{}) {
		log.Fatal("[config] already initialized")
		return &Env{}, errors.New("[config] intialized already")
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("[config] error when loading .env file")
		return &Env{}, errors.New("[config] error when loading .env file")
	}

	log.Print("[config] Initialized .env file")

	setNewEnvironment()

	return &environment, nil
}

func Get() (*Env, error) {
	if environment == (Env{}) {
		return &Env{}, errors.New("[config] not initialized")
	}

	return &environment, nil
}
