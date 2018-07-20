package main

import (
	"testing"
)

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't gott an error but wanted one")
	}
	if got.Error() != want {
		t.Errorf("got '%s', want '%s", got, want)
	}
}

func TestWallet(t *testing.T) {

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, "cannot withdraw, Withdraw insufficient funds")
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
