package wallets

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
	
		got := wallet.Balance()
		expected := Bitcoin(10)
	
		if got != expected {
			t.Errorf("got %s but expected %d", got, expected)
		}
	})

	t.Run("Withdraw", func(t *testing.T){

		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		expected := Bitcoin(10)

		if got != expected {
			t.Errorf("got %s but expected %s", got, expected)
		}

	})

}