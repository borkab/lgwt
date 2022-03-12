package main

import "fmt"

func main() {
	name := "Borka"
	fmt.Println(Hello(name))
}

func Hello(name string) string {
	return "Hello, " + name
}
