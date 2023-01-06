package collection

type Queue[T any] struct {
	inner []T
}

func CreateQueue[T any]() *Queue[T] {
	return &Queue[T]{make([]T, 0)}
}

func (q *Queue[T]) Push(element T) *Queue[T] {
	q.inner = append(q.inner, element)

	return q
}

func (q *Queue[T]) Pop() T {
	elem := q.inner[0]
	q.inner = q.inner[1:]

	return elem
}

func (q *Queue[T]) Length() int {
	return len(q.inner)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.Length() == 0
}
