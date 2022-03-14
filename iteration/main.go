package main

import "fmt"

//const repeatCount = 5

func main() {

	fmt.Println(Repeat("b"))
}

func Repeat(character string) string {
	var repeated string
	var input int
	fmt.Scanln(&input)

	for i := 0; i < input; i++ {
		repeated += character
	}
	return repeated
}
