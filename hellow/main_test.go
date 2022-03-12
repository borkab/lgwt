package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %#q want %#q", got, want)
		}
		// we have refactored our assertion in a function.
		//testing.TB is an interface which *testing.T and *testing.B both satisfys
		//so u can call helper functions from a test or a benchmark
		//t.Helper() is needed to telll thetest suite that this method is a helper.

	}

	t.Run("to a person", func(t *testing.T) {
		got := Hello("Borka", "")
		want := "Hello, Borka"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
