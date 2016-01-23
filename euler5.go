package main

func init() {
	solutions[5] = euler5
}

func euler5() {
	input := make([]uint64, 0, 32)
	for i := uint64(2); i < 21; i++ {
		input = append(input, i)
	}

	println(Lcm(input))
}
