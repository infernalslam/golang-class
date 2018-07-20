package main

import (
	"errors"
	"fmt"
)

type Stringer interface {
	String() string
}
type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Deposit fucntion
func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in test is %v", &w.balance)
	w.balance += amount
}

// Balance function
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw functions
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cannot withdraw, Withdraw insufficient funds")
	}
	w.balance -= amount
	return nil
}
