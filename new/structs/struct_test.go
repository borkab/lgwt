package main

import "testing"

func TestPerimeter(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Width: 10.0, Height: 10.0}
		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %2.f, want %2.f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		got := circle.Perimeter()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %2.f, want %2.f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{Width: 12.0, Height: 6.0}
		got := rectangle.Area()
		want := 72.0

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{Radius: 10}
		got := circle.Area()
		want := 62.83185307179586

		if got != want {
			t.Errorf("got %g, want %g", got, want)
		}
	})
}
