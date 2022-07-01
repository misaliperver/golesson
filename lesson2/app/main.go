package main

import (
	"log"

	"github.com/misaliperver/golesson/lesson2/bootstrap"
	"github.com/misaliperver/golesson/lesson2/config"
	"github.com/misaliperver/golesson/lesson2/models/task"
)

func main() {
	env, err := config.Get()

	// models.Create(models.Task{
	// 	Title:       "Task 1",
	// 	Code:        "TASK1",
	// 	Description: "This is the first task",
	// 	Completed:   false,
	// })

	// models.FindOne("62ae3c50985aac035b65f2d2")

	finded, _ := task.FindAll()

	if len(finded) > 0 {
		log.Println("finded to many object in collection len:", len(finded), ". Not generate new data")
	} else {
		bootstrap.RandomInsert()
	}

	log.Println(env.ConsoleLevel, err)
}
