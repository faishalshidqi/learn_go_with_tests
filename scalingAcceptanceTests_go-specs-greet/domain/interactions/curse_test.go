package interactions_test

import (
	"go-specs-greet/domain/interactions"
	"go-specs-greet/specifications"
	"testing"
)

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(
		t,
		specifications.CurseAdapter(interactions.Curse),
	)
}
