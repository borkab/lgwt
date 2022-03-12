package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Borka")
	want := "Hello, Borka"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
