package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// We're using fmt.Fprint which takes an io.Writer
// (like *bytes.Buffer) and sends a string to it.
func Countdown(out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
