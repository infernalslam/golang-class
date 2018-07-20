package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type SpySleeper struct {
	Calls int
}

type Sleeper interface {
	Sleep()
}

func (s *SpySleeper) Sleep() {
	s.Calls++
	time.Sleep(1 * time.Second)
}

// Countdown test
func Countdown(w io.Writer, s *SpySleeper) {
	for index := 3; 0 <= index; index-- {
		if index == 0 {
			fmt.Fprintf(w, "Go!")
		} else {
			fmt.Fprintf(w, "%v\n", index)
		}
		s.Sleep()
	}
}
func main() {
	Countdown(os.Stdout, &SpySleeper{})
}
