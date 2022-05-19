package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	//	fmt.Printf("address of ballance in test is %v \n", &wallet.balance)
	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}