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
		shape Shapes2D
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
		shape Shapes2D
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
func TestVolume(t *testing.T) {
	checkVolume := func(t testing.TB, shape Shapes3D, want float64) {
		t.Helper()
		got := shape.Volume()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("triangular Pyramids", func(t *testing.T) {
		tri := Triangle{12, 6, 17}
		triPyramid := Pyramid{BaseArea: Triangle.Area(tri), BasePerimeter: Triangle.Perimeter(tri), Height: 21}
		checkVolume(t, triPyramid, 251.99999999999747)
	})

	t.Run("rectangular Pyramids", func(t *testing.T) {
		rec := Rectangle{12, 6}
		recPyramid := Pyramid{BaseArea: Rectangle.Area(rec), BasePerimeter: Rectangle.Perimeter(rec), Height: 21}
		checkVolume(t, recPyramid, 503.99999999999494)
	})

	t.Run("Spheres", func(t *testing.T) {
		sphere := Sphere{Circle{Radius: 10}}
		checkVolume(t, sphere, 4188.790204786381)
	})

	t.Run("Prisms", func(t *testing.T) {
		prism := Prism{Length: 9, Rectangle: Rectangle{Width: 12, Height: 6}}
		checkVolume(t, prism, 648)

	})

}

func TestSurfaceArea(t *testing.T) {

	checkSurfaceArea := func(t testing.TB, shape Shapes3D, want float64) {
		t.Helper()
		got := shape.SurfaceArea()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("triangular Pyramids", func(t *testing.T) {
		tri := Triangle{12, 6, 17}
		triPyramid := Pyramid{BaseArea: Triangle.Area(tri), BasePerimeter: Triangle.Perimeter(tri), SlantHeight: 23}
		checkSurfaceArea(t, triPyramid, 438.5)
	})

	t.Run("rectangular Pyramids", func(t *testing.T) {
		rec := Rectangle{12, 6}
		recPyramid := Pyramid{BaseArea: Rectangle.Area(rec), BasePerimeter: Rectangle.Perimeter(rec), SlantHeight: 23}
		checkSurfaceArea(t, recPyramid, 486)
	})

	t.Run("Spheres", func(t *testing.T) {
		sphere := Sphere{Circle{Radius: 10}}
		checkSurfaceArea(t, sphere, 1256.6370614359173)
	})

	t.Run("Prisms", func(t *testing.T) {
		prism := Prism{Length: 9, Rectangle: Rectangle{Width: 12, Height: 6}}
		checkSurfaceArea(t, prism, 468)
	})
}
