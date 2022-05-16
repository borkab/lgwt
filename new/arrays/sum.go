package main

func Sum(numbers [5]int) int {
	sum := 0
	var num int
	for _, num = range numbers {
		sum += num
	}
	return sum
}
