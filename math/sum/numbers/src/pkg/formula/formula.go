package formula

import (
	"fmt"
	"io"

	"github.com/gookit/color"
)

type Input struct {
	FirstNumber  int
	SecondNumber int
}

func (h Input) Run(writer io.Writer) {
	var result string
	result += color.FgGreen.Render(fmt.Sprintf("The sum of %v and %v is %v\n", h.FirstNumber, h.SecondNumber, h.FirstNumber+h.SecondNumber))

	if _, err := fmt.Fprintf(writer, result); err != nil {
		panic(err)
	}
}
