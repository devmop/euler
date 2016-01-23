package main

func PrimeFactors(number uint64) []uint64 {
	results := make([]uint64, 0, 16)

	primes := Primes(number)

	for _, prime := range primes {
		if number%prime == 0 {
			results = append(results, prime)
		}
	}

	if len(results) == 0 {
		return append(results, number)
	}

	return results
}

//Primes returns all prime factor candidates for n. i.e all primes x such that x * x < ns
func Primes(n uint64) []uint64 {
	if n < 2 {
		return nil
	}
	primes := make([]uint64, 1, 1024)
	primes[0] = 2

	for x := uint64(3); x*x <= n; x += 2 {
		prime := true
		for i := 1; i < len(primes) && primes[i]*primes[i] <= x; i++ {
			if x%primes[i] == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes = append(primes, x)
		}
	}

	return primes
}
