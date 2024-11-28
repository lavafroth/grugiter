package main

type Iterator[T any] interface {
	next() T
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

func main() {
	fib := Fibonacci{0, 1, 0}
	for i := 0; i < 20; i++ {
		f := fib.next()
		println(f)
	}
}
