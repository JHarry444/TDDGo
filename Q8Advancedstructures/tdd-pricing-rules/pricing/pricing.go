package pricing

type Item struct {
	Name  string
	Price int
	Qty   int
}

func Subtotal(items []Item) int {
	s := 0
	for _, it := range items {
		s += it.Price * it.Qty
	}
	return s
}

type DiscountRule func(items []Item, subtotal int) int

func Checkout(items []Item, rules ...DiscountRule) (subtotal, discount, total int) {
	sub := Subtotal(items)
	disc := 0
	for _, rule := range rules {
		disc += rule(items, sub)
	}
	if disc < 0 { disc = 0 }
	if disc > sub { disc = sub }
	return sub, disc, sub - disc
}

func NewPercentOffAbove(threshold int, rate float64) DiscountRule {
	return func(items []Item, subtotal int) int {
		if subtotal >= threshold {
			d := int(float64(subtotal)*rate + 0.5)
			if d < 0 { return 0 }
			return d
		}
		return 0
	}
}

func NewBOGO(itemName string) DiscountRule {
	return func(items []Item, subtotal int) int {
		disc := 0
		for _, it := range items {
			if it.Name == itemName && it.Qty > 1 {
				free := it.Qty / 2
				disc += free * it.Price
			}
		}
		return disc
	}
}

func FilterItems(items []Item, pred func(Item) bool) []Item {
	out := make([]Item, 0, len(items))
	for _, it := range items {
		if pred(it) {
			out = append(out, it)
		}
	}
	return out
}
