package mock

type Topic struct {
	Name   string
	Scores []int
}

type TopicProvider interface {
	GetTopics() []Topic
}

type TopicManager struct {
	Provider TopicProvider
}

func (tm *TopicManager) FindHighestPerTopic() map[string]int {
	results := make(map[string]int)
	topics := tm.Provider.GetTopics()
	for _, topic := range topics {
		if len(topic.Scores) == 0 {
			continue
		}
		highest := topic.Scores[0]
		for _, score := range topic.Scores {
			if score > highest {
				highest = score
			}
		}
		results[topic.Name] = highest
	}
	return results
}
