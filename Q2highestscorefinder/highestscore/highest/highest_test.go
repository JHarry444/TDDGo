package highest

import (
	"testing"
)

func TestFindHighest(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		want    int
		wantErr bool
	}{
		{"typical case", []int{4, 5, -8, 3, 11, -21, 6}, 11, false},
		{"single element", []int{9}, 9, false},
		{"two elements ascending", []int{7, 13}, 13, false},
		{"two elements descending", []int{13, 4}, 13, false},
		{"duplicate highest", []int{3, 5, 5, 2}, 5, false},
		{"empty array", []int{}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FindHighest(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("expected error: %v, got: %v", tt.wantErr, err)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("FindHighest(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
