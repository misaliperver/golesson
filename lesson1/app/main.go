package main

import "github.com/misaliperver/golesson/lesson1/config"

func main() {
	env, err := config.Initialize()

	// env2, err := config.Get()

	// print(env2)
	// print(err)

	// now do something with s3 or whatever
}
