package itertool

type FilterIterator[T any] struct {
	iterable Iterable[T]
	lambda   func(t *T) bool
}

func Filter[T any](iterable Iterable[T], lambda func(t *T) bool) *FilterIterator[T] {
	return &FilterIterator[T]{
		iterable: iterable,
		lambda:   lambda,
	}
}

func (f *FilterIterator[T]) Next() *Optional[T] {
	var o *Optional[T]
	for f.iterable.Next().Export(&o) {
		if !f.lambda(o.GetPointer()) {
			continue
		}
		return o
	}
	return Null[T]()
}
