package topic

import "testing"

func TestFindHighestPerTopic(t *testing.T) {
	input := []Topic{
		{"Physics", []int{56, 67, 45, 89}},
		{"Art", []int{87, 66, 78}},
		{"Comp Sci", []int{45, 88, 97, 56}},
	}
	expected := []Result{
		{"Physics", 89},
		{"Art", 87},
		{"Comp Sci", 97},
	}

	got, err := FindHighestPerTopic(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(got) != len(expected) {
		t.Fatalf("expected %d results, got %d", len(expected), len(got))
	}
	for i, result := range got {
		if result != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], result)
		}
	}
}