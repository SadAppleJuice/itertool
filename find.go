package itertool

func Find[T any](iterable Iterable[T], lambda func(element *T) bool) *T {
	var o *Optional[T]
	for iterable.Next().Export(&o) {
		if lambda(o.GetPointer()) {
			return o.GetPointer()
		}
	}
	return nil
}
