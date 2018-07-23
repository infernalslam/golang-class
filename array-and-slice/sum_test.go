package main

import "testing"

func TestSum(t *testing.T) {

	assertError := func(t *testing.T, got int, want int, numbers []int) {
		t.Helper()
		if want != got {
			t.Errorf("got %d want %d griven, %v", got, want, numbers)
		}
	}

	t.Run("sum numbers array 12345;", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertError(t, got, want, numbers)
	})

	t.Run("sum numbers array 123;", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		assertError(t, got, want, numbers)
	})

}
