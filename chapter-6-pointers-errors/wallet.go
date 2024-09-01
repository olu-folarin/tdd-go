package wallet

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}


func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p and balance is %d \n", &w.balance, w.balance)
	w.balance += amount
	fmt.Printf("address of balance in Deposit is %p and balance is %d \n", &w.balance, w.balance)
}

func (w *Wallet) Total() Bitcoin {
	return w.balance
}