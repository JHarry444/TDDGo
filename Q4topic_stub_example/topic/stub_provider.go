package topic

type StubTopicProvider struct{}

func (s StubTopicProvider) GetTopics() []Topic {
	return []Topic{
		{"Physics", []int{56, 67, 45, 89}},
		{"Art", []int{87, 66, 78}},
	}
}
