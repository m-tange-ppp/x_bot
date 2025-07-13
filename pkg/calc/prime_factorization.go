package calc

func PrimeFactorize(number int) []int {
	factors := []int{}
	targetPrime := 2

	if number < 2 {
		return factors
	}

	for number != 1 {
		for number%targetPrime == 0 {
			number /= targetPrime
			factors = append(factors, targetPrime)
		}
		targetPrime += 1
	}

	return factors
}
