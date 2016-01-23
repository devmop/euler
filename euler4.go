package main

import "fmt"

func init() {
	solutions[4] = euler4
}

func euler4() {
	var result uint64

	for i := uint64(999); i > 99; i-- {
		for j := uint64(999); j > 99; j-- {
			candidate := i * j
			if result < candidate && isPalindrome(candidate) {
				result = candidate
			}
		}
	}

	println(result)
}

func isPalindrome(n uint64) bool {
	val := []rune(fmt.Sprintf("%d", n))
	midpoint := len(val) / 2
	for i := 0; i <= midpoint; i++ {
		if val[i] != val[len(val)-1-i] {
			return false
		}
	}

	return true
}
