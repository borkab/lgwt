package main

import (
	"bytes"
	"testing"
)

//the Buffer type of the bytes package implements the Writer interface,
//because it has the method Write(p []byte) (n int, err error).

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Chris")

	got := buffer.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
