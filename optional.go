package itertool

type Optional[T any] struct {
	value *T
}

func OfPointer[T any](value *T) *Optional[T] {
	return &Optional[T]{
		value: value,
	}
}

func Of[T any](value T) *Optional[T] {
	return &Optional[T]{
		value: &value,
	}
}

func Null[T any]() *Optional[T] {
	return &Optional[T]{
		value: nil,
	}
}

func (o Optional[T]) IsNil() bool {
	return o.value == nil
}

func (o *Optional[T]) GetPointer() *T {
	return o.value
}

func (o *Optional[T]) Get() T {
	return *o.value
}

func (o *Optional[T]) Export(out **Optional[T]) bool {
	if o.IsNil() {
		return false
	}
	*out = o
	return true
}
