package formula

import (
	"errors"
	"fmt"
	"github.com/gookit/color"
	"io"
	"io/ioutil"
	"net/http"
)

type Formula struct {
	CEP string
}

func (h Formula) Run(writer io.Writer) {
	var result string
	address, err := h.RequestAddress()
	if err != nil {
		panic(err)
	}
	result += color.FgGreen.Render(fmt.Sprintf("%s\n", string(address)))
	if _, err := fmt.Fprintf(writer, result); err != nil {
		panic(err)
	}
}

func (h Formula) RequestAddress() ([]byte, error) {
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", h.CEP)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, errors.New("address not found")
	}

	return body, nil
}
