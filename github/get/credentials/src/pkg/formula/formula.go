package formula

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/gookit/color"
)

type Input struct {
	Username string
	Token    string
}

func (h Input) Run(writer io.Writer) {
	var result string
	b, err := json.Marshal(h)
	if err != nil {
		panic(err)
	}
	result += color.FgGreen.Render(string(b) + "\n")

	if _, err := fmt.Fprintf(writer, result); err != nil {
		panic(err)
	}
}
