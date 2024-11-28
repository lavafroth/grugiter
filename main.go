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

func (f *Fibonacci) Next() int {
	defer func() { f.Index += 1 }()
	if f.Index < 2 {
		return 0
	}
	prev := f.Value
	f.Value += f.Previous
	f.Previous = prev
	return f.Value
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
		fmt.Printf("%d\n", f)
	}
}
