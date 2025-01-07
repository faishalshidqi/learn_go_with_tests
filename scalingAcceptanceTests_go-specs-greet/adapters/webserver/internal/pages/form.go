package pages

import (
	"fmt"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

type Form struct {
	Page *rod.Page
}

func (f Form) Greet(name string) error {
	greeting, err := f.Page.Element("#greet-input")
	if err != nil {
		return fmt.Errorf("couldn't find #greet-input on Page")
	}
	return greeting.MustInput(name).Type(input.Enter)
}

func (f Form) Curse(name string) error {
	curse, err := f.Page.Element("#curse-input")
	if err != nil {
		return fmt.Errorf("couldn't find #curse-input on Page")
	}
	return curse.MustInput(name).Type(input.Enter)
}
