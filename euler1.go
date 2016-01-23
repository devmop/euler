package main

func init() {
	solutions[1] = euler1
}

func euler1() {
	total := 0

	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}

	println(total)
}
