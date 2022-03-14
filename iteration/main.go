package main

import (
	"fmt"
	"strings"
)

//const repeatCount = 5

func main() {
	fmt.Println("type a number")
	var input int
	fmt.Scanln(&input)
	chars := "b"

	fmt.Println(Repeat(chars, input))

	fmt.Println("type in a world")
	var in string
	fmt.Scanln(&in)

	fmt.Println("type a character to count")
	var char string
	fmt.Scanln(&char)

	fmt.Println(count(in, char))
}

func Repeat(character string, count int) string {

	var repeated string

	for i := 0; i < count; i++ {
		repeated += character
	}
	return repeated
}

func count(s, ch string) int {

	return strings.Count(s, ch)
}
