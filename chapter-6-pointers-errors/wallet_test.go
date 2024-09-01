package wallet

import (
	"fmt"
	"testing"
)

// run type conversion from int to Bitcoin
func TestWallet(t *testing.T) {
	wallet := Wallet{}
	t.Run("Deposit", func(t *testing.T) {
		wallet.Deposit(Bitcoin(77))
		got := wallet.Total()
		fmt.Printf("address od balance in test is %p. it contains %d btc \n", &wallet.balance, wallet.balance)
		want := Bitcoin(77)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet.Withdraw(Bitcoin(9))
		got := wallet.Total()
		want := Bitcoin(68)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
