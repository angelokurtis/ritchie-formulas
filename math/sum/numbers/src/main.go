package main

import (
	"formula/pkg/formula"
	"os"
	"strconv"
)

func main() {
	input2, err := strconv.Atoi(os.Getenv("NUMBER_ONE"))
	if err != nil {
		panic(err)
	}
	input3, err := strconv.Atoi(os.Getenv("NUMBER_TWO"))
	if err != nil {
		panic(err)
	}

	formula.Input{
		FirstNumber:  input2,
		SecondNumber: input3,
	}.Run(os.Stdout)
}
