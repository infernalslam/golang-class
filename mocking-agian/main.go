package main

import (
	"fmt"
	"io"
)

// Countdown is function loopback count
func Countdown(word io.Writer) {
	fmt.Fprint(word, "3")
}

func main() {
	fmt.Println("Hello, Gopher~!☺️ I'm countdown......")
	// buffer := &bytes.Buffer{}
	// fmt.Printf("result %b", buffer) // &{[] 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 0}%
	// Countdown()
}
