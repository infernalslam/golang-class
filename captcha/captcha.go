package main

import (
	"fmt"
	"math/rand"
	"time"
)

var timeJa int64 = 10

// Random must random
func Random(min, max int) int {
	timeJa = timeJa + 1
	rand.Seed(time.Now().Unix() + timeJa)
	return rand.Intn(max-min) + min
}

// GenerateCaptcha must be return values
func GenerateCaptcha() (string, int) {
	firstCaptcha := Random(10, 99)
	secondCaptcha := Random(10, 99)

	result := 0
	operator := ""
	switch operationRandom := Random(0, 3); operationRandom {
	case 0:
		result = firstCaptcha + secondCaptcha
		operator = "+"
	case 1:
		result = firstCaptcha - secondCaptcha
		operator = "-"
	case 2:
		result = firstCaptcha * secondCaptcha
		operator = "*"
	default:
		result = firstCaptcha / secondCaptcha
		operator = "/"
	}

	captchaString := fmt.Sprintf("%d %s %d", firstCaptcha, operator, secondCaptcha)
	return captchaString, result
}

func main() {
	fmt.Println("Hello Test")
}
