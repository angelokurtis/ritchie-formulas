package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	input1 := os.Getenv("GIT_USER")
	input2 := os.Getenv("GIT_TOKEN")

	formula.Input{
		Username: input1,
		Token:    input2,
	}.Run(os.Stdout)
}
