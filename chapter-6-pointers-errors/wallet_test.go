package wallet

import ("testing"
		"fmt"
	)

	// run type conversion from int to Bitcoin
func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(77))
	got := wallet.Total()
	fmt.Printf("address od balance in test is %p. it contains %d btc \n", &wallet.balance, wallet.balance)
	want := Bitcoin(77)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}