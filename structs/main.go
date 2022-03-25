package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Rectangle struct { //we made a struct so we can take it as a single argument in a func
	Width  float64
	Height float64
}

func main() {
	rect := Rectangle{10.0, 12.5}
	rad := Circle{7.23}
	fmt.Println(Perimeter(rect))
	fmt.Println(rad.Area())
	fmt.Println(rect.Area())
}

func Perimeter(rectangle Rectangle) float64 { //rectangle is a variable type Rectangle what is our struct
	return 2 * (rectangle.Width + rectangle.Height) //so we can call the stuff from our struct
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
