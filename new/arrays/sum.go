package main

func Sum(numbers []int) int {
	sum := 0
	var number int
	for _, number = range numbers {
		sum += number
	}
	return sum
}
