package topic

import (
	"os"
	"testing"
)

func TestWriteResultsToFile(t *testing.T) {
	results := []Result{
		{"Math", 99},
		{"Science", 88},
	}
	filename := "test_output.json"

	err := WriteResultsToFile(results, filename)
	if err != nil {
		t.Fatalf("unexpected error writing file: %v", err)
	}

	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		t.Errorf("expected file %s to be created", filename)
	}

	// Clean up
	os.Remove(filename)
}