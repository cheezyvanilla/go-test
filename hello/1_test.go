package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Bro", "")
		want := "Hello, Bro"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("Say 'hello' in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("In France", func(t *testing.T) {
		got := Hello("Ziath", "French")
		want := "Bonjour, Ziath"

		assertCorrectMessage(t, got, want)
	})
}
