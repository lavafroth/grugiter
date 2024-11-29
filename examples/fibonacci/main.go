package main

import (
	"fmt"
	grug "github.com/lavafroth/grugiter/grug"
)

type Fibonacci struct {
	Index    int
	Value    int
	Previous int
}

// Always return a pointer to your object.
// The iterator follows this pattern.
// If the iterator returns nil, stop iterating.
func (f *Fibonacci) Next() *int {
	defer func() { f.Index += 1 }()
	if f.Index < 2 {
		ret := 0
		return &ret
	}
	prev := f.Value
	f.Value += f.Previous
	f.Previous = prev
	return &f.Value
}

func timesTwo(x int) int {
	return x * 2
}

func main() {
	fib := Fibonacci{0, 1, 0}
	fibIterator := grug.NewIterator(&fib)
	timesTwoMapper := grug.NewMapper(fibIterator, timesTwo)
	plusOneMapper := grug.NewMapper(timesTwoMapper, func(x int) int { return x + 1 })
	divisibleBy3Filter := grug.NewFilter(plusOneMapper, func(x int) bool { return x%3 == 0 })
	for i := 0; i < 20; i++ {
		f := divisibleBy3Filter.Next()
		// We have a guarantee of never receiving a nil value
		// f is safe to dereference
		fmt.Printf("%d\n", *f)
	}
}
