package itertool

func ForEach[T any](iterable Iterable[T], lambda func(element *T)) {
	var o *Optional[T]
	for iterable.Next().Export(&o) {
		lambda(o.GetPointer())
	}
}
