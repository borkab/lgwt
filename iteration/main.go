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

	fmt.Println("type an other character to check")
	var this string
	fmt.Scanln(&this)

	fmt.Println(contain(in, this))

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

func contain(ss, c string) bool {
	return strings.ContainsAny(ss, c)
}
