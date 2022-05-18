package main

import (
	"math"
)

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
	Width  float64
}
type TriangularPyramid struct {
	BaseArea      float64
	BasePerimeter float64
	SlantHeight   float64
}
type Shapes2D interface {
	Area() float64
	Perimeter() float64
}
type Shapes3D interface {
	Volume() float64
	SurfaceArea() float64
}
type Shape interface {
	Shapes2D
	Shapes3D
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (c Circle) Perimeter() float64 {
	return (c.Radius * 2) * math.Pi
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.Height + t.Width
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (tp TriangularPyramid) SurfaceArea() float64 {
	return tp.BaseArea + (tp.BasePerimeter*tp.SlantHeight)*0.5
}
