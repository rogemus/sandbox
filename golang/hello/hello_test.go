package main

import "testing"

func TestHello(t *testing.T) {

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Kacper", "")
		want := "Hello Kacper"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello World' when empty name", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"

		assertCorrectMessage(t, got, want)
	})

  t.Run("say hello in Spanish", func(t *testing.T) {
    got := Hello("Elodie", "Spanish")
    want := "Hola Elodie"
    assertCorrectMessage(t, got, want)
  })
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
