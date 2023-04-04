package itertool

func Collect[T any](iterable Iterable[T]) []T {
	s := make([]T, 0)
	var o *Optional[T]
	for iterable.Next().Export(&o) {
		s = append(s, o.Get())
	}
	return s
}

func CollectN[T any](iterable Iterable[T], n int) []T {
	s := make([]T, 0)
	for i := 0; i < n; i++ {
		o := iterable.Next()
		if o.IsNil() {
			break
		}
		s = append(s, o.Get())
	}
	return s
}
