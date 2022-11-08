package main

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// Stringer interface lets you define how your type is
// printed when used with the %s format string in prints.
type Stringer interface {
	String() string
}

// the w is copy of whatever we called the method from
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
	//unneccessary to dereference the pointer int the function
	//because struct pointers are automatically dereferenced in Go
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
