package main

type PrimesGenerator struct {
	primes    chan uint64
	terminate chan struct{}
}

func PrimeFactors(number uint64) []uint64 {
	results := make([]uint64, 0, 16)

	primes := Primes()
	defer primes.Close()

	for prime := primes.Next(); prime*prime <= number; prime = primes.Next() {
		if number%prime == 0 {
			results = append(results, prime)
		}
	}

	if len(results) == 0 {
		return append(results, number)
	}

	return results
}

func Primes() *PrimesGenerator {
	primes := &PrimesGenerator{make(chan uint64), make(chan struct{})}
	go calculatePrimes(primes.primes, primes.terminate)
	return primes
}

func (p *PrimesGenerator) Next() uint64 {
	return <-p.primes
}

func (p *PrimesGenerator) Close() {
	p.terminate <- struct{}{}
}

func calculatePrimes(ch chan uint64, terminate chan struct{}) {
	ch <- 2

	primes := make([]uint64, 1, 1024)
	primes[0] = 2

	for x := uint64(3); ; x += 2 {
		prime := true
		for i := 1; i < len(primes) && primes[i]*primes[i] <= x; i++ {
			if x%primes[i] == 0 {
				prime = false
				break
			}
		}

		if prime {
			primes = append(primes, x)
			select {
			case ch <- x:
			case <-terminate:
				close(ch)
				return
			}
		}
	}
}
