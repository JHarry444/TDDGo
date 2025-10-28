package topic

import "errors"

// FindHighest returns the highest integer from a non-empty array of integers.
func FindHighest(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, errors.New("input array is empty")
	}

	max := numbers[0]
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}
	return max, nil
}
