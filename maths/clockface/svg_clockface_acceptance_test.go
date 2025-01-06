package clockface

import (
	"bytes"
	"encoding/xml"
	"reflect"
	"testing"
	"time"
)

// SVG was generated 2024-12-26 18:22:38 by https://xml-to-go.github.io/ in Ukraine.
type SVG struct {
	XMLName xml.Name `xml:"svg"`
	//Text    string   `xml:",chardata"`
	Xmlns   string `xml:"xmlns,attr"`
	Width   string `xml:"width,attr"`
	Height  string `xml:"height,attr"`
	ViewBox string `xml:"viewBox,attr"`
	Version string `xml:"version,attr"`
	Circle  Circle `xml:"circle"`
	Line    []Line `xml:"line"`
}

type Circle struct {
	//Text  string `xml:",chardata"`
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
	//Style string `xml:"style,attr"`
}

type Line struct {
	//Text  string `xml:",chardata"`
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
	//Style string `xml:"style,attr"`
}

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	b := bytes.Buffer{}
	SVGWriter(&b, tm)

	svg := SVG{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := 150.0
	y2 := 60.0

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
	}

	t.Errorf("expected to find the second hand with x2 of %+v and y2 of %+v, in the SVG output \"%v\"", x2, y2, b.String())
}

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 60}},
		{simpleTime(0, 0, 30), Line{150, 150, 150, 240}},
	}
	for _, c := range cases {
		t.Run(testFormat(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)
			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(c.line, svg.Line) {
				t.Errorf("expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterMinuteHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(0, 0, 0), Line{150, 150, 150, 70}},
	}
	for _, c := range cases {
		t.Run(testFormat(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)
			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(c.line, svg.Line) {
				t.Errorf("expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func TestSVGWriterHourHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{simpleTime(6, 0, 0), Line{150, 150, 150, 200}},
	}
	for _, c := range cases {
		t.Run(testFormat(c.time), func(t *testing.T) {
			b := bytes.Buffer{}
			SVGWriter(&b, c.time)
			svg := SVG{}
			xml.Unmarshal(b.Bytes(), &svg)
			if !containsLine(c.line, svg.Line) {
				t.Errorf("expected to find the hour hand line %+v, in the SVG line %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(l Line, ls []Line) bool {
	for _, line := range ls {
		if reflect.DeepEqual(line, l) {
			return true
		}
	}
	return false
}

/*

func TestSecondHandAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)
	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(bytes.Buffer{}, tm)
	assertEqual(t, want, got)
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)
	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(tm)
	assertEqual(t, want, got)
}
*/

func assertEqual(t *testing.T, want, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
