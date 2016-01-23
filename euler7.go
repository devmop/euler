package main

func init() {
	solutions[7] = euler7
}

func euler7() {
	p := Primes()
	defer p.Close()

	for i := 0; i < 10000; i++ {
		p.Next()
	}

	println(p.Next())
}
