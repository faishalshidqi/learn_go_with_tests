package pages

import (
	"fmt"
	"github.com/go-rod/rod"
)

type Reply struct {
	Page *rod.Page
}

func (r Reply) ReadReply() (string, error) {
	reply, err := r.Page.Element("#reply")
	if err != nil {
		return "", fmt.Errorf("couldn't find #reply on Page")
	}
	return reply.Text()
}
