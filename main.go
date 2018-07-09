package main

import(
	"fmt"
)

func main() {
	// prime := isPrime(9)
	// fmt.Printf("%t\n", prime)
	number := []int {2, 4, 5}
	calculate(number)
	// fmt.Printf("%v\n", number)
	// send value array in function calculate

}


func calculate(number []int) {
	fmt.Printf("%v\n", number)
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
