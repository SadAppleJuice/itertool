package itertool

type EnumerateIterator[T any] struct {
	iterable Iterable[T]
	index    int
}

type Index[T any] struct {
	Index int
	Value *T
}

func Enumerate[T any](iterable Iterable[T]) *EnumerateIterator[T] {
	return &EnumerateIterator[T]{
		iterable: iterable,
		index:    0,
	}
}

func (e *EnumerateIterator[T]) Next() *Optional[Index[T]] {
	o := e.iterable.Next()
	if o.IsNil() {
		return Null[Index[T]]()
	}
	e.index += 1
	return Of(Index[T]{
		Index: e.index - 1,
		Value: o.GetPointer(),
	})
}
