package main

import "fmt"

//const repeatCount = 5

func main() {
	var input int
	fmt.Scanln(&input)
	char := "b"

	fmt.Println(Repeat(char, input))
}

func Repeat(character string, count int) string {

	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}
