package main

import "testing"

func TestPerimeter(t *testing.T) {

	/*
		checkPerimeter := func(t testing.TB, shape Shape, want float64) {
			t.Helper()
			got := shape.Perimeter()

			if got != want {
				t.Errorf("got %g, want %g", got, want)
			}
		}

		t.Run("rectangles", func(t *testing.T) {
			rectangle := Rectangle{Width: 10.0, Height: 10.0}
			checkPerimeter(t, rectangle, 40.0)
		})

		t.Run("circles", func(t *testing.T) {
			circle := Circle{Radius: 10.0}
			checkPerimeter(t, circle, 62.83185307179586)

		})
	*/
	perimeterTests := []struct { //we are creating a slice of struct
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12, 6}, want: 36.0},
		{shape: Circle{10}, want: 62.83185307179586},
		{shape: Triangle{12, 6, 17}, want: 35.0},
	}

	for _, tt := range perimeterTests { //then we iterate over the slice
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	/*
		checkArea := func(t testing.TB, shape Shape, want float64) {
			t.Helper()
			got := shape.Area()

			if got != want {
				t.Errorf("got %g, want %g", got, want)
			}
		}

		t.Run("rectangles", func(t *testing.T) {
			rectangle := Rectangle{Width: 12.0, Height: 6.0}
			checkArea(t, rectangle, 72.0)

		})

		t.Run("circles", func(t *testing.T) {
			circle := Circle{Radius: 10}
			checkArea(t, circle, 314.1592653589793)
		})
	*/

	areaTests := []struct { //we are creating a slice of struct
		shape Shape
		want  float64
	}{
		{shape: Rectangle{12, 6}, want: 72.0},
		{shape: Circle{10}, want: 314.1592653589793},
		{shape: Triangle{12, 6, 17}, want: 36.0},
	}

	for _, tt := range areaTests { //then we iterate over the slice
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
