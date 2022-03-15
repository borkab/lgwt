package main

import "fmt"

func main() {
	numb := [5]int{1, 2, 3, 4, 5}
	fmt.Println(sum(numb))
}

func sum(numbers [5]int) int {
	sum := 0
	for i := 0; i < 5; i++ {
		sum += numbers[i]
	}
	return sum
}
