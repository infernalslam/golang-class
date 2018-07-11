package main

import "fmt"

type stack []int

func main() {
	var value stack
	// value variable send to function push()
	value.push(1)
	value.push(2)
	value.push(3)
	value.push(4)
	value.push(5)

	// value send to pop() function
	value.pop()
	value.pop()
	fmt.Print(value)
}

func (s *stack) push(number int) stack {
	// fmt.Print(a) //5
	*s = append(*s, number)
	return *s
}

func (s *stack) pop() (int, error) {
	// fmt.Print(len(*s)) // Get length of a pointer array in golang
	*s = (*s)[:len(*s)-1]
	return 0, nil
}
