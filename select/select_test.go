package _select

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest", func(t *testing.T) {
		slowServer := makeDelayedServer(50 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()

		slowurl := slowServer.URL
		fasturl := fastServer.URL
		want := fasturl
		got, err := ConfigurableRacer(slowurl, fasturl, 10*time.Millisecond)
		assertNoError(t, err)
		assertEqual(t, want, got)
	})
	t.Run("returns an error if a server doesn't respond within 10ms", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Millisecond)
		serverB := makeDelayedServer(12 * time.Millisecond)
		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 10*time.Millisecond)
		assertError(t, err)
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	t.Helper()
	if !reflect.DeepEqual(a, b) {
		t.Errorf("want %v got %v", a, b)
	}
}

func assertError(t *testing.T, err error) {
	t.Helper()
	if err == nil {
		t.Error("expected an error but didn't get one")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Error("expected no error but got one")
	}
}
