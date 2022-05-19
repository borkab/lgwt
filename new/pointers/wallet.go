package main

import (
	"errors"
	"fmt"
)

// Bitcoin represent a number of Bitcoins
type Bitcoin int

//Wallet stores the number of Bitcoin someone owns
type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}

//Deposit will add some Bitcoin to a wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	//	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

//Balance returns the number os Bitcoin a wallet has
func (w *Wallet) Balance() Bitcoin {
	return w.balance //the same: (*w).balance   because in Go the
	// structs pointers are automatically dereferenced
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh, no")
	} //errors.New creates a new error with a message of your choosing.

	w.balance -= amount
	return nil
}

func (b *Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
