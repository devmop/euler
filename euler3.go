package main

func init() {
	solutions[3] = euler3
}

func euler3() {
	factors := PrimeFactors(600851475143)

	println(factors[len(factors)-1])
}
