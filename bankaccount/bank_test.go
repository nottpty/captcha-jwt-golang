package backaccount

import "testing"

func TestNewAccount(t *testing.T) {
	newAccount := New("Papaya")
	tempAccount := Account{
		id:      1,
		name:    "Papaya",
		balance: 0,
	}
	if newAccount.id != tempAccount.id {
		t.Errorf("expect %d got %d", newAccount.id, tempAccount.id)
	}
	if newAccount.name != tempAccount.name {
		t.Errorf("expect %s got %s", newAccount.name, tempAccount.name)
	}
	if newAccount.balance != tempAccount.balance {
		t.Errorf("expect %d got %d", newAccount.balance, tempAccount.balance)
	}
}

func TestSaveAccount(t *testing.T) {
	account := New("Alien")
	newData := &Account{
		id:      1,
		name:    "Hellen",
		balance: 2300,
	}
	Save(account)
	if account.id != 1 {
		t.Errorf("expect 1 got %d", account.id)
	}
	if account.name != "Alien" {
		t.Errorf("expect Alien got %s", account.name)
	}
	if account.balance != 0 {
		t.Errorf("expect 0 got %d", account.balance)
	}
	Save(newData)
	if account.id != 1 {
		t.Errorf("expect 1 got %d", account.id)
	}
	if account.name != "Hellen" {
		t.Errorf("expect Hellen got %s", account.name)
	}
	if account.balance != 2300 {
		t.Errorf("expect 2300 got %d", account.balance)
	}
}

func TestFindAccountByName(t *testing.T) {
	if getLengthBankAccount() == 0 {
		account := New("Alien")
		Save(account)
		findAccount1 := FindByName("Alien")
		if findAccount1.id != 1 {
			t.Errorf("expect 1 got %d", findAccount1.id)
		}
		if findAccount1.name != "Alien" {
			t.Errorf("expect Alien got %s", findAccount1.name)
		}
		if findAccount1.balance != 0 {
			t.Errorf("expect 0 got %d", findAccount1.balance)
		}
	} else {
		findAccount1 := FindByName("Hellen")
		if findAccount1.id != 1 {
			t.Errorf("expect 1 got %d", findAccount1.id)
		}
		if findAccount1.name != "Hellen" {
			t.Errorf("expect Hellen got %s", findAccount1.name)
		}
		if findAccount1.balance != 2300 {
			t.Errorf("expect 2300 got %d", findAccount1.balance)
		}
	}

}

func TestFindAccountNotFound(t *testing.T) {
	findAccount := FindByName("Mind")
	if findAccount != nil {
		t.Error("expect nil got", *findAccount)
	}
}
