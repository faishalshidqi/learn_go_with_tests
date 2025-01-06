package clockface

import (
	"fmt"
	"io"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 50
	clockCenterX     = 150
	clockCenterY     = 150
	svgStart         = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
	bezel  = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
	svgEnd = `</svg>`
)

func SVGWriter(w io.Writer, tm time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	SecondHand(w, tm)
	MinuteHand(w, tm)
	HourHand(w, tm)
	io.WriteString(w, svgEnd)
}

func makeHand(point Point, length float64) Point {
	point = Point{point.X * length, point.Y * length}             // scale
	point = Point{point.X, -point.Y}                              // flip
	point = Point{point.X + clockCenterX, point.Y + clockCenterY} // translate
	return point
}

// SecondHand is the unit vector of the second hand
// of an analogue clock at time `t`
// represented as a Point
func SecondHand(w io.Writer, t time.Time) {
	point := makeHand(SecondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, point.X, point.Y)
}

// MinuteHand is the unit vector of the minute hand
// of an analogue clock at time `t`
// represented as a Point
func MinuteHand(w io.Writer, t time.Time) {
	point := makeHand(MinuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, point.X, point.Y)
}

// HourHand is the unit vector of the hour hand
// of an analogue clock at time `t`
// represented as a Point
func HourHand(w io.Writer, t time.Time) {
	point := makeHand(HourHandPoint(t), hourHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, point.X, point.Y)
}
