package main

import (
	"formula/pkg/formula"
	"os"
)

func main() {
	CEP := os.Getenv("CEP")

	formula.Formula{
		CEP: CEP,
	}.Run(os.Stdout)
}
