package main

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 10}
		checkPerimeter(t, rectangle, 40.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{7}
		checkPerimeter(t, circle, 43.982297150257104)
	})

	t.Run("triangles", func(t *testing.T) {
		triangle := Triangle{12, 6, 9}
		checkPerimeter(t, triangle, 27)
	})
}

func TestArea(t *testing.T) {
	//table driven test
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{shape: Triangle{Height: 12, Base: 6, Width: 9}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
	/*	checkArea := func(t testing.TB, shape Shape, want float64) {
			t.Helper()
			got := shape.Area()
			if got != want {
				t.Errorf("got %g want %g", got, want)
			}
		}

		t.Run("rectangles", func(t *testing.T) {

			rectangle := Rectangle{12, 6}
			checkArea(t, rectangle, 72.0)
		})

		t.Run("circles", func(t *testing.T) {

			circle := Circle{10}
			checkArea(t, circle, 314.1592653589793)
		})
	*/

}
