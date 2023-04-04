package iterators

import "github.com/SadAppleJuice/itertool"

type SliceIterator[T any] struct {
	slice *[]T
	index int
}

func NewSliceIterator[T any](slice *[]T) *SliceIterator[T] {
	return &SliceIterator[T]{
		slice: slice,
		index: 0,
	}
}

func (s *SliceIterator[T]) Next() *itertool.Optional[T] {
	if s.index >= len(*s.slice) {
		return itertool.Null[T]()
	}
	s.index += 1
	return itertool.Of((*s.slice)[s.index-1])
}
