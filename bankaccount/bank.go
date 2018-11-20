package backaccount

var bankAccount []*Account
var id int

func init() {
	bankAccount = []*Account{}
	id++
}

// New use for create bank account
func New(name string) *Account {
	tempAccount := &Account{
		id:      id,
		name:    name,
		balance: 0,
	}

	return tempAccount
}

// Save if not found id of account in local storage or update if found id of account in local storage
func Save(account *Account) {
	for _, tempAccount := range bankAccount {
		if tempAccount.id == account.id {
			tempAccount.name = account.name
			tempAccount.balance = account.balance
			return
		}
	}
	bankAccount = append(bankAccount, account)
	id++
}

// FindByName use for find account by using name of account
func FindByName(name string) *Account {
	for _, account := range bankAccount {
		if account.name == name {
			return account
		}
	}
	return nil
}

func getLengthBankAccount() int {
	return len(bankAccount)
}
