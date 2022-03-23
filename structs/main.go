package main

import "fmt"

func main() {
	w := 10.0
	h := 12.5

	fmt.Println(Perimeter(w, h))
}

func Perimeter(width, heigth float64) float64 {
	return 2 * (width + heigth)
}
