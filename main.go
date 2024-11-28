package main

import (
	"fmt"
)

type IteratorImpl[T any] interface {
	next() T
}

type Iterator[T IteratorImpl[V], V any] struct {
	State T
	Next  func() V
}

func NewIterator[T IteratorImpl[V], V any](state T) Iterator[T, V] {
	return Iterator[T, V]{
		State: state,
		Next:  state.next,
	}
}

type Mapper[T IteratorImpl[V], V any, U any] struct {
	State Iterator[T, V]
	MapFn func(V) U
}

func (s Mapper[T, V, U]) next() U {
	return s.MapFn(s.State.Next())
}

type Filter[T IteratorImpl[V], V any] struct {
	State Iterator[T, V]
	MapFn func(V) bool
}

func (s Filter[T, V]) next() V {
	nextElement := s.State.Next()
	for !s.MapFn(nextElement) {
		nextElement = s.State.Next()
	}
	return nextElement
}

func NewMapper[T IteratorImpl[V], V any, U any](iterator Iterator[T, V], mapFn func(V) U) Iterator[Mapper[T, V, U], U] {
	mapper := Mapper[T, V, U]{
		State: iterator,
		MapFn: mapFn,
	}
	return Iterator[Mapper[T, V, U], U]{
		State: mapper,
		Next:  mapper.next,
	}
}

func NewFilter[T IteratorImpl[V], V any](iterator Iterator[T, V], mapFn func(V) bool) Iterator[Filter[T, V], V] {
	filter := Filter[T, V]{
		State: iterator,
		MapFn: mapFn,
	}
	return Iterator[Filter[T, V], V]{
		State: filter,
		Next:  filter.next,
	}
}

// That was all for defining my iterator.
// This is where you start implementing your program.

type Fibonacci struct {
	Index    int
	Value    int
	Previous int
}

func (f *Fibonacci) next() int {
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
	fibIterator := NewIterator(&fib)
	timesTwoMapper := NewMapper(fibIterator, timesTwo)
	plusOneMapper := NewMapper(timesTwoMapper, func(x int) int { return x + 1 })
	divisibleBy3Filter := NewFilter(plusOneMapper, func(x int) bool { return x%3 == 0 })
	for i := 0; i < 20; i++ {
		f := divisibleBy3Filter.Next()
		fmt.Printf("%d\n", f)
	}
}
