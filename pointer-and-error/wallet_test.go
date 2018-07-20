package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		// fmt.Printf("address of balance in test is %v", &wallet.balance)

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d,", got, want)
		}
	})
}
