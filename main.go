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
	mapper := NewMapper(fibIterator, timesTwo)
	for i := 0; i < 20; i++ {
		f := mapper.Next()
		fmt.Printf("%d\n", f)
	}
}
