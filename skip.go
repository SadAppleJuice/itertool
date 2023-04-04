package itertool

func Skip[T any](iterable Iterable[T], count int) Iterable[T] {
	for i := 0; i < count; i++ {
		iterable.Next()
	}
	return iterable
}
