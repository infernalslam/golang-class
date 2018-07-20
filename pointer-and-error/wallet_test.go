package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Error("wanted an error but didn't get one")
		}
	}

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err)
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})

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

	// t.Run("withdraw", func(t *testing.T) {
	// 	wallet := Wallet{balance: Bitcoin(20)}
	// 	wallet.Withdraw(Bitcoin(10))
	// 	got := wallet.Balance()
	// 	want := Bitcoin(10)

	// 	if got != want {
	// 		t.Errorf("got %d want %d,", got, want)
	// 	}
	// })
}
