package topic

import (
	"testing"

)

func TestFindHighestPerTopicWithStub(t *testing.T) {
	provider := StubTopicProvider{}
	got, err := FindHighestPerTopicFromProvider(provider)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := []Result{
		{"Physics", 89},
		{"Art", 87},
	}

	if len(got) != len(want) {
		t.Fatalf("expected %d results, got %d", len(want), len(got))
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %v, want %v", got[i], want[i])
		}
	}
}
