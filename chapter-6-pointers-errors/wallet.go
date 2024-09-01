package wallet

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

type Stringer interface {
	String() string
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %p and balance is %d \n", &w.balance, w.balance)
	w.balance += amount
	fmt.Printf("address of balance in Deposit is %p and balance is %d \n", &w.balance, w.balance)
}

func (w *Wallet) Total() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) {
	fmt.Printf("address of balance in Withdraw is %p and balance before withdrawal is %d \n", &w.balance, w.balance)
	w.balance -= amount
	fmt.Printf("address of balance in Withdraw is %p and new balance is %d \n", &w.balance, w.balance)
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
