package main

import (
	"fmt"
)

type stack []int

func main() {
	var value stack
	// value.push(5)
	// value.pop()
	fmt.Print("%v\n", value)
}

func (s stack) push(a int) stack {
	return s
}

func (s stack) pop() (int, error) {
	return 0, nil
}
