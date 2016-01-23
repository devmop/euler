package main

func Lcm(input []uint64) uint64 {
	result := Product(input)

	for len(input) > 0 {
		candidate := result / input[0]

		success := true

		for _, v := range input {
			if candidate%v != 0 {
				success = false
				break
			}
		}

		if success {
			result = candidate
		} else {
			input = append(input[:0], input[1:]...)
		}
	}

	return result
}

func Product(input []uint64) uint64 {
	result := uint64(1)

	for _, v := range input {
		result *= v
	}

	return result
}
