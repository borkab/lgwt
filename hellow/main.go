package main

import "fmt"

const french = "Frernch"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjur, "

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchHelloPrefix + name
	}
	return englishHelloPrefix + name
}
