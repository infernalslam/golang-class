package main

import (
	"fmt"
	"io"
)

// Greet function
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}
