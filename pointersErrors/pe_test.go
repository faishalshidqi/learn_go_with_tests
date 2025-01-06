package pointersErrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		assertResult(t, got, want)
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(100)}
		err := wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(90)
		assertResult(t, got, want)
		assertNoError(t, err)
	})
	t.Run("withdrawing insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(100)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(110))
		assertResult(t, wallet.balance, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertResult(t *testing.T, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if err.Error() != want {
		t.Errorf("got \"%s\" want \"%s\"", err, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
