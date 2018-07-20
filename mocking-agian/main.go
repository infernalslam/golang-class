package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const countdownStart = 3

// Countdown is function loopback count
func Countdown(word io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(word, i)
	}
	fmt.Fprint(word, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
