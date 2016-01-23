package main

type FibGenerator struct {
	current, next uint64
}

func (f *FibGenerator) Next() uint64 {
	f.current, f.next = f.next, f.current+f.next
	return f.current
}

func Fibs() *FibGenerator {
	return &FibGenerator{0, 1}
}
