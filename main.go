package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3}
	result, average := calculate(numbers)
	fmt.Printf("sum = %d, avg = %f\n", result, average)
	// day2
	var slice = []string{"a", "b", "c", "d"}
	copySlice := append([]string(nil), slice...)
	fmt.Println(copySlice)
	slice = append(slice, "b", "o", "y")
	fmt.Println(slice)

	sliceNumber := []int{1, 2, 4, 5, 7, 8}
	reverseNumber := reverseSlice(sliceNumber)
	fmt.Println(reverseNumber)
	// fmt.Println(slice[:2])
	// fmt.Println(slice[1:2])
	// fmt.Println(slice[4:])
}

func calculate(numbers []int) (int, float32) {
	result := 0
	count := 0
	for _, number := range numbers {
		if isPrime(number) {
			result += number
			count++
		}
	}
	return result, float32(result) / float32(count)
}

func isPrime(num int) bool {
	prime := num != 1
	for i := 2; i < num; i++ {
		if num%i == 0 {
			prime = false
			break
		}
	}
	return prime
}

func reverseSlice(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		tmp := nums[i]
		nums[i] = nums[j]
		nums[j] = tmp
	}
	return nums
}
