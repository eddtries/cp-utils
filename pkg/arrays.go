package cputils

func sum(numbers []int) (accumulator int) {
	for _, number := range numbers {
		accumulator += number
	}

	return
}
