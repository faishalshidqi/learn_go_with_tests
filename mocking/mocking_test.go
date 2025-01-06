package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type SpyCountdownOperations struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("slept for 3 seconds", func(t *testing.T) {
		spySleeper := &SpySleeper{}

		buffer := &bytes.Buffer{}
		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		assertEqual(t, got, want)
		assertEqual(t, spySleeper.Calls, 3)
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		assertEqual(t, want, spySleepPrinter.Calls)
	})
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	assertEqual(t, spyTime.durationSlept, sleepTime)
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Errorf("got %v want %v", a, b)
	}
}
