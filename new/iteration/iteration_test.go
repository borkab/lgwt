package iteration

import "testing"

func TestRepeat(t *testing.T) {
	character := "a"
	repeated := Repeat(character)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
