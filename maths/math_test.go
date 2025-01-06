package maths

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, want, got interface{}) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want: %v, got: %v", want, got)
	}
}
