package pointers

import "testing"

func TestWallet(t *testing.T) {
	checker := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("want %s got %s", want, got)
		}

	}

	errChecker := func(t *testing.T, err error, want string) {
		t.Helper()

		if err == nil {
			t.Fatal("wanted error but got none")
		}

		if err.Error() != want {
			t.Errorf("got %q want %q", err, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		checker(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{20}

		wallet.Withdraw(Bitcoin(10))
		checker(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		checker(t, wallet, startingBalance)
		errChecker(t, err, "balance is not enough")
	})
}
