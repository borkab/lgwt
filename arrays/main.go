package main

import "fmt"

func main() {
	numb := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(numb))
}

func sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
