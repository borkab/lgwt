package main

func Sum(numbers []int) int {
	sum := 0
	var number int
	for _, number = range numbers {
		sum += number
	}
	return sum
}

func SumAll(slice1 []int, slice2 []int) []int {
	var all []int
	s1 := 0
	var n1 int
	for _, n1 = range slice1 {
		s1 += n1
	}
	s2 := 0
	var n2 int
	for _, n2 = range slice2 {
		s2 += n2
	}
	all = append(all, s1, s2)
	return all
}
