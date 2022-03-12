package main

import (
	"fmt"
)

const croatian = "Croatian"
const hungarian = "Hungarian"
const french = "Frernch"
const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjur, "
const hungarianHelloPrefix = "Szia, "
const croatianHelloPrefix = "Bok, "

func main() {
	fmt.Println(Hello("world", ""))
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greretingPrefix(language) + name
}

func greretingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix

	case french:
		prefix = frenchHelloPrefix

	case hungarian:
		prefix = hungarianHelloPrefix

	case croatian:
		prefix = croatianHelloPrefix

	default:
		prefix = englishHelloPrefix
	}
	return
}
