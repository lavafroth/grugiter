package grugiter

type IteratorImpl[V any] interface {
	Next() *V
}

type Iterator[T IteratorImpl[V], V any] struct {
	State T
	Next  func() *V
}

func NewIterator[T IteratorImpl[V], V any](state T) Iterator[T, V] {
	return Iterator[T, V]{
		State: state,
		Next:  state.Next,
	}
}

type Mapper[T IteratorImpl[V], V any, U any] struct {
	State Iterator[T, V]
	MapFn func(V) U
}

func (s Mapper[T, V, U]) Next() *U {
	someOrNone := s.State.Next()
	if someOrNone == nil {
		return nil
	}
	mapped := s.MapFn(*someOrNone)
	return &mapped
}

type Filter[T IteratorImpl[V], V any] struct {
	State Iterator[T, V]
	MapFn func(V) bool
}

func (s Filter[T, V]) Next() *V {
	nextElement := s.State.Next()
	if nextElement == nil {
		return nil
	}
	for !s.MapFn(*nextElement) {
		nextElement = s.State.Next()
		if nextElement == nil {
			return nil
		}
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
		Next:  mapper.Next,
	}
}

func NewFilter[T IteratorImpl[V], V any](iterator Iterator[T, V], mapFn func(V) bool) Iterator[Filter[T, V], V] {
	filter := Filter[T, V]{
		State: iterator,
		MapFn: mapFn,
	}
	return Iterator[Filter[T, V], V]{
		State: filter,
		Next:  filter.Next,
	}
}
