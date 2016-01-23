package main

func init() {
	solutions[10] = euler10
}

func euler10() {
	p := Primes()
	defer p.Close()

	total := uint64(0)

	for prime := p.Next(); prime < 2000000; prime = p.Next() {
		total += prime
	}

	println(total)
}
