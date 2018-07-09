package main

import(
	"fmt"
)

func main() {
	// prime := isPrime(9)
	// fmt.Printf("%t\n", prime)
	numbers := []int {2, 4, 5}
	calculate(numbers)
	// fmt.Printf("%v\n", number)
	// send value array in function calculate

}


func calculate(numbers []int) {
	// fmt.Printf("%v\n", number)
	// something check for range
	for _, number := range numbers {
		fmt.Printf("%v\n", number)
	}
}

func isPrime(num int) bool {
	numberIsPrime := true
	for i := 2; i <= num / 2; i++ {
		if num % i == 0 {
			numberIsPrime = false
			break
		}
	}
	return numberIsPrime
}
