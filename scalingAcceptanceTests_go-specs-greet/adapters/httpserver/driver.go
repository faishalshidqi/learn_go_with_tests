package httpserver

import (
	"fmt"
	"io"
	"net/http"
)

type Driver struct {
	BaseURL string
	Client  *http.Client
}

func (d *Driver) Greet(name string) (string, error) {
	res, err := http.Get(fmt.Sprintf(d.BaseURL+"/greet?name=%s", name))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func (d *Driver) Curse(name string) (string, error) {
	res, err := http.Get(fmt.Sprintf(d.BaseURL+"/curse?name=%s", name))
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
