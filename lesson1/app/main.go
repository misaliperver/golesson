package main

import (
	"github.com/misaliperver/golesson/lesson1/config"
	"github.com/misaliperver/golesson/lesson1/models"
)

func init() {
	config.Initialize()
}

func main() {
	// env, err := config.Get()

	models.Create("Kitap Al", false)
}
