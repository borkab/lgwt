package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//func Greet(writer *bytes.Buffer, name string) {

func Greet(writer io.Writer, name string) {
	//fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)

	//fmt.Printf() prints our data to stdout.
	//fmt.Fprintf() takes a Writer to send the string to

}
func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
