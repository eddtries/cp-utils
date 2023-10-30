package cputils

import "errors"

func CollatzSequence(n int) ([]int, error) {
	if n <= 0 {
		return []int{}, errors.New("non-positive integer")
	}
	sequence := []int{n}
	for {
		if n == 1 {
			return sequence, nil
		}
		if n%2 == 0 {
			n /= 2
			sequence = append(sequence, n)
		} else {
			n = 3*n + 1
			sequence = append(sequence, n)
		}
	}
}
