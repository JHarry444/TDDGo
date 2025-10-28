package main

import (
	"fmt"
	"log"

	"example.com/topic/topic"
)

func main() {
	topics := []topic.Topic{
		{"Physics", []int{56, 67, 45, 89}},
		{"Art", []int{87, 66, 78}},
		{"Comp Sci", []int{45, 88, 97, 56}},
	}

	results, err := topic.FindHighestPerTopic(topics)
	if err != nil {
		log.Fatal(err)
	}

	err = topic.WriteResultsToFile(results, "top_scores.json")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Scores written to top_scores.json")
}
