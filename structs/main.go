package main

import "github.com/xoreo/go-basics/structs/types"

func main() {
	myFroggy, err := types.NewFrog(
		"big hunk",
		700,
		true,
	)

	stupidFrog, err := types.NewFrog(
		"I can't ree because im stupid",
		5,
		false,
	)

	if err != nil {
		panic(err)
	}

	err = myFroggy.Ree()
	if err != nil {
		panic(err)
	}

	err = stupidFrog.Ree()
	if err != nil {
		panic(err)
	}
}
