package helloWorld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("IQ", "")
		want := "Hello, IQ!"
		assertMessage(t, got, want)
	})
	t.Run("saying 'Hello, World!; when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World!"
		assertMessage(t, got, want)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Juan", "spanish")
		want := "Hola, Juan!"
		assertMessage(t, got, want)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Louis", "french")
		want := "Bonjour, Louis!"
		assertMessage(t, got, want)
	})
	t.Run("in Bahasa Indonesia", func(t *testing.T) {
		got := Hello("Malih", "indonesian")
		want := "Halo, Malih!"
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
