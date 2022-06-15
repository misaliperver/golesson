package main

import "github.com/misaliperver/golesson/lesson1/config"

func main() {
	env, err := config.Initialize()
	println(err)
	println(env.ConsoleLevel)

}
