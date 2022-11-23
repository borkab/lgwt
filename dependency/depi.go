package main

import (
	"fmt"
	"io"
	"os"
)

//func Greet(writer *bytes.Buffer, name string) {

func Greet(writer io.Writer, name string) {
	//fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)

	//fmt.Printf() prints our data to stdout.
	//fmt.Fprintf() takes a Writer to send the string to

}

func main() {
	Greet(os.Stdout, "Mila")
}
