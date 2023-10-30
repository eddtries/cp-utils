package cputils

func IsPrime(x int) bool {
	if x <= 1 {
		return false
	}
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}

	return true
}

func PrimesGenerator() func() int {
	currentPrime := 1
	return func() int {
		i := currentPrime + 1

		for {
			if IsPrime(i) {
				currentPrime = i
				return currentPrime
			}
			i++
		}
	}
}
