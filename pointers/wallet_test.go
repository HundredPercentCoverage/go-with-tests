package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	t.Run("deposit", func(t *testing.T) {
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		err := wallet.Withdraw(10)
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(0))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		err := wallet.Withdraw(10)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(0))
	})
}

// Put helpers here so basic assertions can be read first
// then these if necessary

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func assertError(t testing.TB, err, want error) {
	t.Helper()
	if err == nil {
		// Must use Fatal here to end tests
		// Subsequent use of err.Error() would be on a null pointer
		// and cause a panic
		t.Fatal("expected an error, didn't get one")
	}

	if err != want {
		t.Errorf("got %q, wanted %q", err, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error("expected no error but got one")
	}
}
