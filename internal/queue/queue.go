package queue

type Queue[T any] struct {
	slice []T
}

func New[T any](es ...T) Queue[T] {
	var q Queue[T]
	q.slice = es
	return q
}

func (q Queue[T]) Len() int { return len(q.slice) }

func (q Queue[T]) Peek() (e T, ok bool) {
	if q.Len() == 0 {
		return e, false
	}
	return q.slice[0], true
}

func (q *Queue[T]) Enqueue(v T) {
	q.slice = append(q.slice, v)
}

func (q *Queue[T]) Dequeue() (e T, ok bool) {
	if q.Len() == 0 {
		return e, false
	}
	e, q.slice = q.slice[0], q.slice[1:]
	return e, true
}
