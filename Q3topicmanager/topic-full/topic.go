package topic

import (
	"errors"
)

type Topic struct {
	Name   string
	Scores []int
}

type Result struct {
	Name  string
	Score int
}

// FindHighestPerTopic returns the highest score per topic.
func FindHighestPerTopic(topics []Topic) ([]Result, error) {
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

		results = append(results, Result{Name: t.Name, Score: highest})
	}
	return results, nil
}
