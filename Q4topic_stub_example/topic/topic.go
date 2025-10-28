package topic

type Topic struct {
	Name   string
	Scores []int
}

type Result struct {
	Name  string
	Score int
}

type TopicProvider interface {
	GetTopics() []Topic
}

func FindHighestPerTopicFromProvider(provider TopicProvider) ([]Result, error) {
	topics := provider.GetTopics()
	results := []Result{}

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
		results = append(results, Result{Name: topic.Name, Score: highest})
	}
	return results, nil
}
