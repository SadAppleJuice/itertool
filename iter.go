package itertool

type Iterable[T any] interface {
	Next() *Optional[T]
}
