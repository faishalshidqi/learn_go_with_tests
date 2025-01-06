package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"
	assertResult(t, repeated, expected)
}

func BenchmarkRepeat(b *testing.B) {
	Repeat("a", b.N)
}

func ExampleRepeat() {
	repeated := Repeat("oi", 5)
	fmt.Println(repeated)
	// Output: oioioioioi
}

func assertResult(t *testing.T, repeated, expected string) {
	if repeated != expected {
		t.Errorf("Got %s, expected %s", repeated, expected)
	}
}
