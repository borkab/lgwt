package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	name := "Borka"
	fmt.Println(Hello(name))
}

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
