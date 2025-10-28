package topic

import "testing"

func TestFindHighest(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
		wantErr  bool
	}{
		{"normal case", []int{4, 5, -8, 3, 11, -21, 6}, 11, false},
		{"empty array", []int{}, 0, true},
		{"single element", []int{5}, 5, false},
		{"duplicate highest", []int{7, 7, 3}, 7, false},
		{"two elements", []int{7, 13}, 13, false},
		{"reverse two", []int{13, 4}, 13, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindHighest(tt.input)
			if (err != nil) != tt.wantErr {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}