package itertool

type MapIterator[T any, O any] struct {
	iterable Iterable[T]
	lambda   func(t *T) O
}

func Map[T any, O any](iterable Iterable[T], lambda func(t *T) O) *MapIterator[T, O] {
	return &MapIterator[T, O]{
		iterable: iterable,
		lambda:   lambda,
	}
}

func (m *MapIterator[T, O]) Next() *Optional[O] {
	v := m.iterable.Next()
	if v.IsNil() {
		return Null[O]()
	}
	return Of(m.lambda(v.GetPointer()))
}
