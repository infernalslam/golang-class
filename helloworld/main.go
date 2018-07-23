package main

import "fmt"

func main() {
	fmt.Println(Hello("Gopher"))
}

func Hello(name string) string {
	if name != "" {
		return "Hello, " + name
	}
	return "Hello"
}
