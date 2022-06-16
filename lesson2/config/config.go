package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Db struct {
	Uri  string
	Name string
}

type Env struct {
	DefaultLang   string
	AppListenPort string
	Db            Db
	ConsoleLevel  string
}

var environment *Env

func setNewEnvironment() {
	environment = &Env{
		DefaultLang:   os.Getenv("DEFAULT_LANG"),
		AppListenPort: os.Getenv("APP_LISTEN_PORT"),
		Db: Db{
			Uri:  os.Getenv("MONGO_URI"),
			Name: os.Getenv("DBNAME"),
		},
		ConsoleLevel: os.Getenv("CONSOLE_LEVEL"),
	}
}

func Initialize() (Env, error) {
	if *environment != (Env{}) {
		log.Fatal("[config] already initialized")
		return Env{}, errors.New("[config] intialized already")
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("[config] error when loading .env file")
		return Env{}, errors.New("[config] error when loading .env file")
	}

	log.Print("[config] Initialized .env file")

	setNewEnvironment()

	return *environment, nil
}

func Get() (Env, error) {
	if *environment == (Env{}) {
		panic("[config] not initialized")
	}

	log.Println("[config] Get() successfully")

	return *environment, nil
}
