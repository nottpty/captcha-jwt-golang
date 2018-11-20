package backaccount

// Account for bank
type Account struct {
	id      int
	name    string
	balance int
}

// Withdraw money from account
func (a *Account) Withdraw(amount int) {
	a.balance = a.balance - amount
}

// Deposit money to account
func (a *Account) Deposit(amount int) {
	a.balance = a.balance + amount
}

// Balance of money in account
func (a *Account) Balance() int {
	return a.balance
}
