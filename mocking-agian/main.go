package main

import (
	"bytes"
	"fmt"
)

// w io.Writer

// Countdown is function loopback count
func Countdown(number *bytes.Buffer) {
	fmt.Fprint(number, "3")
}

func main() {
	fmt.Println("Hello, Gopher~!☺️ I'm countdown......")
	// buffer := &bytes.Buffer{}
	// fmt.Printf("result %b", buffer) // &{[] 0 [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] 0}%
	// Countdown()
}
