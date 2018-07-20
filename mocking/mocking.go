package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Countdown test
func Countdown(w io.Writer) {
	for index := 3; 0 <= index; index-- {
		if index == 0 {
			fmt.Fprintf(w, "Go!")
		} else {
			fmt.Fprintf(w, "%v\n", index)
		}
		time.Sleep(1 * time.Second)
	}
}
func main() {
	Countdown(os.Stdout)
}
