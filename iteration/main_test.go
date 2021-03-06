package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	rep := Repeat("foo", 3)
	fmt.Println(rep)
	//output: foofoofoo
}

func TestCount(t *testing.T) {
	got := count("lilla", "l")
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestContain(t *testing.T) {
	got := contain("kukmuki", "u")
	want := true

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
