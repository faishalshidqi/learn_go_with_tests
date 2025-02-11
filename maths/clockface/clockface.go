package clockface

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	secondsInClock     = 2 * secondsInHalfClock
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

// Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X, Y float64
}

func secondsInRadians(t time.Time) float64 {
	//return float64(t.Second()) * (math.Pi / 30)
	return math.Pi / (secondsInHalfClock / (float64(t.Second())))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / hoursInClock) + math.Pi/(hoursInHalfClock/float64(t.Hour()%hoursInClock))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func SecondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func MinuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func HourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}
