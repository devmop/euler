package main

func init() {
	solutions[6] = euler6
}

func euler6() {
	n := uint64(100)

	sum := (n / 2) * (n + 1)

	squares := uint64(0)

	for i := uint64(1); i <= n; i++ {
		squares += (i * i)
	}

	println((sum * sum) - squares)
}
