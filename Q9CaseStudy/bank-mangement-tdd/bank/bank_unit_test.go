package bank_test

import (
	"bank-management-tdd/bank"
	"fmt"
	"testing"
)

func TestNewAccount(t *testing.T) {
	acc := bank.NewAccount(100, nil)
	if acc.Balance() != 100 {
		t.Errorf("expected balance 100, got %d", acc.Balance())
	}
}

func TestDeposit(t *testing.T) {
	acc := bank.NewAccount(50, nil)
	_ = acc.Deposit(30)
	fmt.Print("Checking balance after deposit...\n")
	fmt.Println("Balance:", acc.Balance())
	if acc.Balance() != 80 {
		t.Errorf("expected balance 80, got %d", acc.Balance())
	}

}

func TestWithdraw(t *testing.T) {
	acc := bank.NewAccount(100, nil)
	_ = acc.Withdraw(40)
	fmt.Print("Checking balance after withdrawal...\n")
	fmt.Println("Balance:", acc.Balance())
	if acc.Balance() != 60 {
		t.Errorf("expected balance 60, got %d", acc.Balance())
	}
}

func TestWithdrawInsufficientFunds(t *testing.T) {
	acc := bank.NewAccount(20, nil)
	err := acc.Withdraw(50)
	fmt.Print("Attempting to withdraw with insufficient funds...\n")
	fmt.Print("Account balance:", acc.Balance(), "\n")
	if err != bank.ErrInsufficientFunds {
		t.Errorf("expected ErrInsufficientFunds, got %v", err)
	}
}
