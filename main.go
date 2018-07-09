package main

import(
	"fmt"
)

func main() {
	numbers := []int {1, 2, 3}
	result, average := calculate(numbers)
	fmt.Printf("sum = %d, avg = %f\n", result, average)

}


func calculate(numbers []int) (int, float32){
	result := 0
	count := 0
	for _, number := range numbers {
		if isPrime(number) {
			result += number
			count++
		}
	}
	return result, float32(result / count)
}


func isPrime(num int) bool {
	prime := num != 1
	for i := 2; i < num; i++ {
		if num % i == 0 {
			prime = false
			break
		}
	}
	return prime
}

