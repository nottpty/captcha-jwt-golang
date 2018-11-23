package bankaccount

// Account for bank
type Account struct {
	ID           int
	Name         string
	BalanceMoney int
}

// Withdraw money from account
func (a *Account) Withdraw(amount int) {
	a.BalanceMoney = a.BalanceMoney - amount
}

// Deposit money to account
func (a *Account) Deposit(amount int) {
	a.BalanceMoney = a.BalanceMoney + amount
}

// Balance of money in account
func (a *Account) Balance() int {
	return a.BalanceMoney
}
