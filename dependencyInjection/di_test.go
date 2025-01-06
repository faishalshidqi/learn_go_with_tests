package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Sh")
	result := buffer.String()
	want := "Hello, Sh"
	assertResult(t, result, want)
}

func assertResult(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
