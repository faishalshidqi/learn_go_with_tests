package specifications

import (
	"github.com/alecthomas/assert/v2"
	"testing"
)

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t *testing.T, meany MeanGreeter) {
	got, err := meany.Curse("Gopher")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go to hell, Gopher!")
}
