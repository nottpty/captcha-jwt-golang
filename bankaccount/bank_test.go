package bankaccount

import "testing"

func TestNewAccount(t *testing.T) {
	newAccount := New("Papaya")
	tempAccount := Account{
		ID:           1,
		Name:         "Papaya",
		BalanceMoney: 0,
	}
	if newAccount.ID != tempAccount.ID {
		t.Errorf("expect %d got %d", newAccount.ID, tempAccount.ID)
	}
	if newAccount.Name != tempAccount.Name {
		t.Errorf("expect %s got %s", newAccount.Name, tempAccount.Name)
	}
	if newAccount.BalanceMoney != tempAccount.BalanceMoney {
		t.Errorf("expect %d got %d", newAccount.BalanceMoney, tempAccount.BalanceMoney)
	}
}

func TestSaveAccount(t *testing.T) {
	account := New("Alien")
	newData := &Account{
		ID:           1,
		Name:         "Hellen",
		BalanceMoney: 2300,
	}
	Save(account)
	if account.ID != 1 {
		t.Errorf("expect 1 got %d", account.ID)
	}
	if account.Name != "Alien" {
		t.Errorf("expect Alien got %s", account.Name)
	}
	if account.BalanceMoney != 0 {
		t.Errorf("expect 0 got %d", account.BalanceMoney)
	}
	Save(newData)
	if account.ID != 1 {
		t.Errorf("expect 1 got %d", account.ID)
	}
	if account.Name != "Hellen" {
		t.Errorf("expect Hellen got %s", account.Name)
	}
	if account.BalanceMoney != 2300 {
		t.Errorf("expect 2300 got %d", account.BalanceMoney)
	}
}

func TestFindAccountByName(t *testing.T) {
	if getLengthBankAccount() == 0 {
		account := New("Alien")
		Save(account)
		findAccount1 := FindByName("Alien")
		if findAccount1.ID != 1 {
			t.Errorf("expect 1 got %d", findAccount1.ID)
		}
		if findAccount1.Name != "Alien" {
			t.Errorf("expect Alien got %s", findAccount1.Name)
		}
		if findAccount1.BalanceMoney != 0 {
			t.Errorf("expect 0 got %d", findAccount1.BalanceMoney)
		}
	} else {
		findAccount1 := FindByName("Hellen")
		if findAccount1.ID != 1 {
			t.Errorf("expect 1 got %d", findAccount1.ID)
		}
		if findAccount1.Name != "Hellen" {
			t.Errorf("expect Hellen got %s", findAccount1.Name)
		}
		if findAccount1.BalanceMoney != 2300 {
			t.Errorf("expect 2300 got %d", findAccount1.BalanceMoney)
		}
	}

}

func TestFindAccountNotFound(t *testing.T) {
	findAccount := FindByName("Mind")
	if findAccount != nil {
		t.Error("expect nil got", *findAccount)
	}
}
