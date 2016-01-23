package main

func init() {
	solutions[9] = euler9
}

func euler9() {
	for a := 1; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			for c := b; c < 1000; c++ {
				if (a+b+c) == 1000 && ((a*a)+(b*b)) == (c*c) {
					println(a * b * c)
					return
				}
			}
		}
	}
}
