package bootstrap

import (
	"fmt"
	"strconv"

	helper "github.com/misaliperver/golesson/lesson2/helpers"
	"github.com/misaliperver/golesson/lesson2/models/task"
)

var generatedTaskCount = 2

func generareNewTask(c chan *task.Task) {
	for i := 0; i < generatedTaskCount; i++ {
		c <- &task.Task{
			Title:       "Task " + strconv.Itoa(i),
			Code:        helper.RandomStringRunes(10),
			Description: helper.randStringRunes(10),
			Completed:   helper.randBool(),
		}
	}

	close(c)
}

func RandomInsert() {
	c := make(chan *task.Task)

	go generareNewTask(c)

	for {
		res, ok := <-c
		if ok == false {
			fmt.Println("Channel Close ", ok)
			break
		}

		fmt.Println("Channel Open ", res, ok)

		task.Create(*res)
	}
}
