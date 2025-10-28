package pricing

import (
	"reflect"
	"testing"
)

func TestSubtotal(t *testing.T) {
	items := []Item{{"Apple", 100, 3}, {"Banana", 50, 2}} // 400
	got := Subtotal(items)
	if got != 400 { t.Fatalf("want 400, got %d", got) }
}

func TestCheckout_NoRules(t *testing.T) {
	items := []Item{{"Apple", 100, 3}} // 300
	sub, disc, total := Checkout(items)
	if sub != 300 || disc != 0 || total != 300 {
		t.Fatalf("want (300,0,300), got (%d,%d,%d)", sub, disc, total)
	}
}

func TestCheckout_PercentOffAboveThreshold(t *testing.T) {
	items := []Item{{"Apple", 1000, 3}} // 3000
	rule := NewPercentOffAbove(2000, 0.10)
	sub, disc, total := Checkout(items, rule)
	if sub != 3000 || disc != 300 || total != 2700 {
		t.Fatalf("want (3000,300,2700), got (%d,%d,%d)", sub, disc, total)
	}
}

func TestCheckout_BOGO_ForSpecificItem(t *testing.T) {
	items := []Item{{"Socks", 200, 3}, {"Tee", 1000, 1}}
	rule := NewBOGO("Socks")
	sub, disc, total := Checkout(items, rule)
	if sub != 200*3+1000 || disc != 200 || total != (200*3+1000-200) {
		t.Fatalf("unexpected totals: (%d,%d,%d)", sub, disc, total)
	}
}

func TestCheckout_ComposedRules(t *testing.T) {
	items := []Item{{"Socks", 200, 4}, {"Tee", 1200, 2}} // 3200
	bogo := NewBOGO("Socks")
	percent := NewPercentOffAbove(3000, 0.05) // 160
	sub, disc, total := Checkout(items, bogo, percent)
	if sub != 3200 || disc != 560 || total != 2640 {
		t.Fatalf("want (3200,560,2640), got (%d,%d,%d)", sub, disc, total)
	}
}

func TestFilterItems_LambdaPredicate(t *testing.T) {
	items := []Item{{"Apple", 100, 3}, {"Tee", 1200, 1}, {"Socks", 200, 2}}
	pred := func(it Item) bool { return it.Price >= 500 }
	got := FilterItems(items, pred)
	want := []Item{{"Tee", 1200, 1}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("want %v, got %v", want, got)
	}
}

func TestCheckout_Guards_NoNegativeOrOverDiscount(t *testing.T) {
	items := []Item{{"Cheap", 10, 1}} // subtotal=10
	tooMuch := func(_ []Item, _ int) int { return 9999 }
	negative := func(_ []Item, _ int) int { return -5 }
	sub, disc, total := Checkout(items, tooMuch, negative)
	if sub != 10 || disc != 10 || total != 0 {
		t.Fatalf("want (10,10,0), got (%d,%d,%d)", sub, disc, total)
	}
}
