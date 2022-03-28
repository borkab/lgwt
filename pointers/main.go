package main

type Wallet struct {
	balance int
}

func main() {

}

func (w *Wallet) Deposit(amount int) {
	//fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int { // w is a copy of whatever we called the method from
	return w.balance
}
