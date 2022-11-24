package main

import (
	"fmt"
	"io"
	"os"
)

// We're using fmt.Fprint which takes an io.Writer
// (like *bytes.Buffer) and sends a string to it.
func Countdown(out io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
