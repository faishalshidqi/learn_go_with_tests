package main_test

import (
	"fmt"
	"go-specs-greet/adapters"
	"go-specs-greet/adapters/httpserver"
	"go-specs-greet/specifications"
	"net/http"
	"testing"
	"time"
)

func TestGreeterServer(t *testing.T) {
	var (
		port    = "8080"
		baseURL = fmt.Sprintf("http://localhost:%s", port)
		driver  = httpserver.Driver{BaseURL: baseURL, Client: &http.Client{
			Timeout: 1 * time.Second,
		}}
	)

	adapters.StartDockerServer(t, port, "httpserver")
	specifications.CurseSpecification(t, &driver)
	specifications.GreetSpecification(t, &driver)
}
