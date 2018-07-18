package main

import "testing"

func TestRandom(t *testing.T) {
	total := Random(10, 99)
	if total < 10 || total > 99 {
		t.Errorf("Sum was incorrect, got: %d", total)
	}
}

func TestGenerateCaptcha(t *testing.T) {
	firstCaptcha := Random(10, 99)
	secondCaptcha := Random(10, 99)
	if firstCaptcha < 10 || firstCaptcha > 99 {
		t.Errorf("Sum was incorrect, got: %d", firstCaptcha)
	}
	if secondCaptcha < 10 || secondCaptcha > 99 {
		t.Errorf("Sum was incorrect, got: %d", secondCaptcha)
	}
}
