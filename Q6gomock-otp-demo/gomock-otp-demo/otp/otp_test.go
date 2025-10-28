package otp

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestSendOTP_SuccessStoresOTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockOTPService(ctrl)
	am := NewAuthManager(mockSvc)
	email := "user@example.com"

	mockSvc.EXPECT().SendOTP(email).Return("123456", nil).Times(1)

	if err := am.SendOTP(email); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := am.store[email]; got != "123456" {
		t.Fatalf("expected stored otp 123456, got %s", got)
	}
}

func TestSendOTP_PropagatesError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockOTPService(ctrl)
	am := NewAuthManager(mockSvc)
	email := "user@example.com"

	mockSvc.EXPECT().SendOTP(email).Return("", errors.New("send failed")).Times(1)

	if err := am.SendOTP(email); err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestVerifyOTP_TrueWhenServiceVerifiesAndStoredMatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockOTPService(ctrl)
	am := NewAuthManager(mockSvc)
	email := "user@example.com"
	stored := "654321"

	am.store[email] = stored

	mockSvc.EXPECT().VerifyOTP(email, stored).Return(true).Times(1)

	if ok := am.VerifyOTP(email, stored); !ok {
		t.Fatalf("expected verification true, got false")
	}
}

func TestVerifyOTP_FalseWhenNoStoredOTP(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockOTPService(ctrl)
	am := NewAuthManager(mockSvc)
	email := "nobody@example.com"

	if ok := am.VerifyOTP(email, "whatever"); ok {
		t.Fatalf("expected false, got true")
	}
}

func TestVerifyOTP_FalseWhenStoredMismatch(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSvc := NewMockOTPService(ctrl)
	am := NewAuthManager(mockSvc)
	email := "user@example.com"

	am.store[email] = "111111" 
	if ok := am.VerifyOTP(email, "222222"); ok {
		t.Fatalf("expected false due to mismatch, got true")
	}
}
