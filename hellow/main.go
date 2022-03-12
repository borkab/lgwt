package main

import "fmt"

const englishHelloPrefix = "Hello, "

func main() {
	name := "Borka"
	fmt.Println(Hello(name))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return englishHelloPrefix + name
}
