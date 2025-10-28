package tagparser

import "strings"

// ParseTags splits a comma-delimited string and returns a cleaned list of tags.
func ParseTags(input string) []string {
	rawTags := strings.Split(input, ",")
	var result []string

	for _, raw := range rawTags {
		tag := strings.TrimSpace(raw)
		if tag != "" {
			result = append(result, tag)
		}
	}

	return result
}
