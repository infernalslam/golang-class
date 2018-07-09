package main

import(
	"fmt"
)

func main() {
	prime := isPrime(9)
	fmt.Printf("%t\n", prime)
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
