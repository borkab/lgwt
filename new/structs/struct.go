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
type Pyramid struct {
	BaseArea      float64
	BasePerimeter float64
	SlantHeight   float64
	Height        float64
}
type Sphere struct {
	Circle
}

type Prism struct {
	Rectangle Rectangle
	Length    float64
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

func (p Pyramid) SurfaceArea() float64 {
	return p.BaseArea + (p.BasePerimeter*p.SlantHeight)*0.5
}

func (p Pyramid) Volume() float64 {
	return 0.33333333333333 * p.BaseArea * p.Height
}

func (s Sphere) Volume() float64 {
	return 1.33333333333333 * math.Pi * s.Radius * s.Radius * s.Radius
}
func (s Sphere) SurfaceArea() float64 {
	return 4 * math.Pi * s.Radius * s.Radius
}
func (p Prism) SurfaceArea() float64 {
	return 2 * (p.Length*p.Rectangle.Width + p.Length*p.Rectangle.Height + p.Rectangle.Height*p.Rectangle.Width)
}
func (p Prism) Volume() float64 {
	return p.Length * p.Rectangle.Height * p.Rectangle.Width
}
