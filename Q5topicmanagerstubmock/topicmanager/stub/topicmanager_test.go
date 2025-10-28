package stub

import (
	"reflect"
	"testing"
)

type StubTopicProvider struct{}

func (s StubTopicProvider) GetTopics() []Topic {
	return []Topic{
		{"Math", []int{40, 70, 90}},
		{"Science", []int{60, 80, 100}},
	}
}

func TestFindHighestPerTopic_WithStub(t *testing.T) {
	provider := StubTopicProvider{}
	tm := TopicManager{Provider: provider}
	expected := map[string]int{
		"Math":    90,
		"Science": 100,
	}
	actual := tm.FindHighestPerTopic()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
