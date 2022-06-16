package main

import (
	config "github.com/misaliperver/golesson/lesson1/config"
	db "github.com/misaliperver/golesson/lesson1/db"
	models "github.com/misaliperver/golesson/lesson1/models"
)

func init() {
	config.Initialize()
	db.ConnectDB()
}

func main() {
	// env, err := config.Get()

	models.Create("Kitap Al", false)
}
