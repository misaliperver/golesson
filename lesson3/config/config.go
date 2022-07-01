package config

import (
	"errors"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Db struct {
	Uri  string
	Name string
}

type Env struct {
	DefaultLang   string
	AppListenPort string
	AppPrefork    bool
	Db            Db
	ConsoleLevel  string
	LoggingFormat string
}

var environment *Env

func setNewEnvironment() {
	appPrefork, _ := strconv.ParseBool(os.Getenv("APP_PREFORK"))

	environment = &Env{
		DefaultLang:   os.Getenv("DEFAULT_LANG"),
		AppListenPort: os.Getenv("APP_LISTEN_PORT"),
		AppPrefork:    appPrefork,
		Db: Db{
			Uri:  os.Getenv("MONGO_URI"),
			Name: os.Getenv("DBNAME"),
		},
		ConsoleLevel:  os.Getenv("CONSOLE_LEVEL"),
		LoggingFormat: os.Getenv("LOGGING_FORMAT"),
	}
}

func Initialize() (*Env, error) {
	if environment != nil {
		log.Fatal("[config] already initialized")
		return nil, errors.New("[config] intialized already")
	}

	err := godotenv.Load()

	if err != nil {
		log.Fatal("[config] error when loading .env file", err)
		return nil, errors.New("[config] error when loading .env file")
	}

	log.Print("[config] Initialized .env file")

	setNewEnvironment()

	return environment, nil
}

func Get() (*Env, error) {
	if environment == nil {
		panic("[config] not initialized")
	}

	log.Println("[config] Get() successfully")

	return environment, nil
}

func init() {
	println("[config] init()")
	Initialize()
}
