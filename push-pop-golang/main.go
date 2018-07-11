package main

import "fmt"

type stack []int

func main() {
	var value stack
	// value variable send to function push()
	value.push(5)
	value.push(6)
	fmt.Print(value)
}

func (s *stack) push(number int) stack {
	// fmt.Print(a) //5
	*s = append(*s, number)
	return *s
}

func (s stack) pop() (int, error) {
	return 0, nil
}
