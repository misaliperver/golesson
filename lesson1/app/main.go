package main

import "github.com/misaliperver/golesson/lesson1/config"

func main() {
	// env, err := config.Initialize()
	// print(err)
	// print(env.consoleLevel)

	cos := config.Env{consoleLevel: "debug"}

	print(cos.consoleLevel)

	// now do something with s3 or whatever
}
