package topic

import (
	"encoding/json"
	"os"
)

// WriteResultsToFile writes the topic-score results to a JSON file.
func WriteResultsToFile(results []Result, filename string) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}
