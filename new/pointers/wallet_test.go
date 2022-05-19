package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		//	fmt.Printf("address of ballance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		//	fmt.Printf("address of ballance in test is %v \n", &wallet.balance)
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
