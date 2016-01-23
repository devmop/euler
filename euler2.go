package main

func init() {
	solutions[2] = euler2
}

func euler2() {
	var total uint64

	fibs := Fibs()

	for i := fibs.Next(); i < 4000000; i = fibs.Next() {
		if (i % 2) == 0 {
			total += i
		}
	}

	println(total)
}
