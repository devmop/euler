package main

import (
	"flag"
	"fmt"
	"strconv"
)

var solutions = make([]func(), 100)

func main() {
	flag.Parse()

	for _, v := range flag.Args() {
		problem, err := strconv.Atoi(v)
		if err == nil {
			if 1 > problem || problem >= len(solutions) || solutions[problem] == nil {
				fmt.Printf("No solution for euler problem %d\n", problem)
			} else {
				fmt.Printf("Solution for euler problem %d \n", problem)
				solutions[problem]()
			}
		} else {
			fmt.Printf("Invalid value %s got error '%s' \n", v, err.Error())
		}
	}
}
