package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

// struc zone
type Sleeper interface {
	Sleep()
}

// Countdown is function loopback count
func Countdown(word io.Writer) {
	for i := countdownStart; i > 0; i-- {
		time.Sleep(1 * time.Second)
		fmt.Fprintln(word, i)
	}
	time.Sleep(1 * time.Second)
	fmt.Fprint(word, finalWord)
}

func main() {
	Countdown(os.Stdout)
}
