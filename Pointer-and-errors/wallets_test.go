package wallets

import (
	"testing"
)


func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, w Wallet, want Bitcoin){
		t.Helper()
		got := w.Balance()
		if got != want { 
			t.Errorf("got %s but expected %d", got, want)
		}
	}
	
	assertError := func(t testing.TB, err error){
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	}
	
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, 10)
	})

	t.Run("Withdraw", func(t *testing.T){
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, 10)

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T){
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err)
		assertBalance(t, wallet, startingBalance)

	})

}