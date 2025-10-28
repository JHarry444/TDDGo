package bank

import "errors"

var ErrInsufficientFunds = errors.New("insufficient funds")

type Logger interface {
	Log(message string) error
}

type Account struct {
	balance int
	logger  Logger
}

func NewAccount(initial int, logger Logger) *Account {
	return &Account{balance: initial, logger: logger}
}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) Deposit(amount int) error {
	a.balance += amount
	if a.logger != nil {
		return a.logger.Log("deposit completed")
	}
	return nil
}

func (a *Account) Withdraw(amount int) error {
	if amount > a.balance {
		return ErrInsufficientFunds
	}
	a.balance -= amount
	if a.logger != nil {
		return a.logger.Log("withdrawal completed")
	}
	return nil
}
