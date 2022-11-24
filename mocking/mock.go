package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// we make a dependency as an interface. this lets us
// then use a real Sleeper in main and a spy Sleeper in our tests.
type Sleeper interface {
	Sleep()
}

// so our Countdown function won't be responsible for how long
// the sleep is.
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// We're using fmt.Fprint which takes an io.Writer
// (like *bytes.Buffer) and sends a string to it.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
