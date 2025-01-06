package interactions_test

import (
	"go-specs-greet/domain/interactions"
	"go-specs-greet/specifications"
	"testing"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(
		t,
		specifications.GreetAdapter(interactions.Greet),
	)
}
