package main

import (
	"fmt"
	"io"
	"net/http"
)

// Greet function
func Greet(w io.Writer, name string) {
	fmt.Fprintf(w, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// Greet(os.Stdout, "Elodie")
	http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler))
}
