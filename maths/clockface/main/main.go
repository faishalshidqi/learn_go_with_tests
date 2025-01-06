package main

import (
	"learn_go_with_tests/maths/clockface"
	"os"
	"time"
)

func main() {
	t := time.Now()
	/*
		b := bytes.Buffer{}
		secondHand := clockface.SecondHand(&b, t)
		io.WriteString(os.Stdout, svgStart)
		io.WriteString(os.Stdout, bezel)
		io.WriteString(os.Stdout, secondHandTag(secondHand))
		io.WriteString(os.Stdout, svgEnd)
	*/
	clockface.SVGWriter(os.Stdout, t)
}
