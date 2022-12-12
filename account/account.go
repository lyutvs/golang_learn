package account

import (
	"errors"
	"fmt"
)

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

func (account *Account) Deposit(amount int) {
	account.balance += amount
}

func (account Account) Balance() int {
	return account.balance
}

func (account Account) WithDraw(amount int) error {
	if account.balance < amount {
		return errNoMoney
	}
	account.balance -= amount

	return nil
}

func (account *Account) ChangeOwner(newOwner string) {
	account.owner = newOwner
}

func (account Account) Owner() string {
	return account.owner
}

func (account Account) String() string {
	return fmt.Sprint(account.Owner(), "'s account. \nHas : ", account.Balance())
}
