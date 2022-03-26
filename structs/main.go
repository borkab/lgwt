package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Rectangle struct { //we made a struct so we can take it as a single argument in a func
	Width  float64
	Height float64
}

func main() {
	rect := Rectangle{10.0, 12.5}
	circ := Circle{7.23}
	tri := Triangle{12, 6}

	fmt.Println(Perimeter(rect))
	fmt.Println(circ.Area())
	fmt.Println(rect.Area())
	fmt.Println(tri.Area())
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

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}
