package main

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	input := fmt.Scanln(&input)
	repeated := Repeat("a")
	expected := "a" * input

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
