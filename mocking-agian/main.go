package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown is function loopback count
func Countdown(word io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(word, i)
	}
	fmt.Fprint(word, "Go!")
}

func main() {
	Countdown(os.Stdout)
}
