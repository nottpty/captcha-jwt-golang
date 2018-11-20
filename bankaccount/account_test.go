package backaccount

import "testing"

func TestBalance(t *testing.T) {
	account := getTempAccount()
	if account.Balance() != 3000 {
		t.Error("Expect 3000 got", account.Balance())
	}
}

func TestWithdraw(t *testing.T) {
	account := getTempAccount()
	account.Deposit(2000)
	if account.Balance() != 5000 {
		t.Error("Expect 5000 got", account.Balance())
	}
}

func TestDeposit(t *testing.T) {
	account := getTempAccount()
	account.Withdraw(1300)
	if account.Balance() != 1700 {
		t.Error("Expect 1700 got", account.Balance())
	}
}

func getTempAccount() Account {
	account := Account{
		id:      0,
		name:    "nameTest",
		balance: 3000,
	}
	return account
}
