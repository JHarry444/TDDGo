package mock

import (
	"reflect"
	"testing"
)

type MockTopicProvider struct {
	called bool
}

func (m *MockTopicProvider) GetTopics() []Topic {
	m.called = true
	return []Topic{
		{"English", []int{55, 65}},
		{"History", []int{45, 85}},
	}
}

func TestFindHighestPerTopic_WithMock(t *testing.T) {
	mock := &MockTopicProvider{}
	tm := TopicManager{Provider: mock}

	results := tm.FindHighestPerTopic()

	if !mock.called {
		t.Error("Expected GetTopics to be called")
	}

	expected := map[string]int{
		"English": 65,
		"History": 85,
	}
	if !reflect.DeepEqual(results, expected) {
		t.Errorf("Expected %v, got %v", expected, results)
	}
}
