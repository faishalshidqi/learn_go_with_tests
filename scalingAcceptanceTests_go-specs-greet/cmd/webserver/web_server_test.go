package main

import (
	"fmt"
	"github.com/alecthomas/assert/v2"
	"go-specs-greet/adapters"
	"go-specs-greet/adapters/webserver"
	"go-specs-greet/specifications"
	"testing"
)

func TestGreeterServer(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	var (
		port            = "8081"
		driver, cleanup = webserver.NewDriver(fmt.Sprintf("http://localhost:%s", port))
	)
	t.Cleanup(func() {
		assert.NoError(t, cleanup())
	})
	adapters.StartDockerServer(t, port, "webserver")
	specifications.GreetSpecification(t, driver)
	specifications.CurseSpecification(t, driver)
}
