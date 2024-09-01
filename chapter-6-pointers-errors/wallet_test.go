package wallet

import (
	"fmt"
	"testing"
)

// run type conversion from int to Bitcoin
func TestWallet(t *testing.T) {
	// deduplicate the code
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Total()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(77))
		fmt.Printf("address od balance in test is %p. it contains %d btc \n", &wallet.balance, wallet.balance)
		assertBalance(t, wallet, Bitcoin(77))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(77)}
		wallet.Withdraw(Bitcoin(9))
		assertBalance(t, wallet, Bitcoin(68))
	})
}
