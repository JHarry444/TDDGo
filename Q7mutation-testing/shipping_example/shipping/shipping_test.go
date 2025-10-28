package shipping

import "testing"

func TestFreeShipping(t *testing.T) {
	got := CalculateShippingCost(1500)
	if got != 0 {
		t.Errorf("Expected 0, got %d", got)
	}
}

func TestStandardShipping(t *testing.T) {
	got := CalculateShippingCost(500)
	if got != 50 {
		t.Errorf("Expected 50, got %d", got)
	}
}

func TestInvalidOrderAmount(t *testing.T) {
	got := CalculateShippingCost(-1)
	if got != -1 {
		t.Errorf("Expected -1 for invalid input, got %d", got)
	}
}

func TestExactlyFreeShippingThreshold(t *testing.T) {
	got := CalculateShippingCost(1000)
	if got != 0 {
		t.Errorf("Expected 0 for order amount exactly 1000, got %d", got)
	}
}

func TestZeroOrderAmount(t *testing.T) {
	got := CalculateShippingCost(0)
	if got != 50 {
		t.Errorf("Expected 50 for order amount 0, got %d", got)
	}
}
