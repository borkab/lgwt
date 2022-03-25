package main

import (
	"fmt"
)

type Rectangle struct { //we made a struct so we can take it as a single argument in a func
	Width  float64
	Height float64
}

func main() {
	rect := Rectangle{10.0, 12.5}

	fmt.Println(Perimeter(rect))
	fmt.Println(Area(rect))
}

func Perimeter(rectangle Rectangle) float64 { //rectangle is a variable type Rectangle what is our struct
	return 2 * (rectangle.Width + rectangle.Height) //so we can call the stuff from our struct
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
