package main

import (
	"fmt"
)

func main() {
	w := 10.0
	h := 12.5

	fmt.Println(Perimeter(w, h))
	fmt.Println(Area(w, h))
}

func Perimeter(width, heigth float64) float64 {
	return 2 * (width + heigth)
}

func Area(width, heigth float64) float64 {
	return width * heigth
}
