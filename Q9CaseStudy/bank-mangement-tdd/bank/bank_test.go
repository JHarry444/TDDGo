package bank_test

import (
	"bank-management-tdd/bank"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestAccountWithLogger(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("Deposit logs transaction and updates balance", func(t *testing.T) {
		mockLogger := bank.NewMockLogger(ctrl)
		mockLogger.EXPECT().Log("deposit completed").Return(nil)

		acc := bank.NewAccount(100, mockLogger)
		err := acc.Deposit(50)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		fmt.Print("Checking balance after deposit...\n")
		fmt.Println("Balance:", acc.Balance())

		if acc.Balance() != 150 {
			t.Errorf("expected balance 150, got %d", acc.Balance())
		}
	})

	t.Run("Withdraw logs transaction and updates balance", func(t *testing.T) {
		mockLogger := bank.NewMockLogger(ctrl)
		mockLogger.EXPECT().Log("withdrawal completed").Return(nil)

		acc := bank.NewAccount(100, mockLogger)
		err := acc.Withdraw(40)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		fmt.Print("Checking balance after withdrawal...\n")
		fmt.Println("Balance:", acc.Balance())

		if acc.Balance() != 60 {
			t.Errorf("expected balance 60, got %d", acc.Balance())
		}
	})

	t.Run("Withdraw fails with insufficient funds", func(t *testing.T) {
		mockLogger := bank.NewMockLogger(ctrl)
		// No log expected since withdrawal should fail

		acc := bank.NewAccount(20, mockLogger)
		err := acc.Withdraw(50)

		fmt.Print("Attempting to withdraw with insufficient funds...\n")
		fmt.Print("Account balance:", acc.Balance(), "\n")
		if err != bank.ErrInsufficientFunds {
			t.Errorf("expected ErrInsufficientFunds, got %v", err)
		}
	})
}
