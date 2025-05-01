package stack

type Stack[T any] struct {
	slice []T
}

func New[T any](es ...T) Stack[T] {
	var s Stack[T]
	s.slice = es
	return s
}

func (s Stack[T]) Len() int { return len(s.slice) }

func (s Stack[T]) last() int { return s.Len() - 1 }

func (s Stack[T]) Peek() (e T, ok bool) {
	if s.Len() == 0 {
		return e, false
	}
	return s.slice[0], true
}

func (s *Stack[T]) Push(v T) {
	s.slice = append(s.slice, v)
}

func (s *Stack[T]) Pop() (e T, ok bool) {
	if s.Len() == 0 {
		return e, false
	}
	e, s.slice = s.slice[s.last()], s.slice[:s.last()]
	return e, true
}

func (s *Stack[T]) PopN(n int) (o []T, ok bool) {
	for range n {
		e, ok := s.Pop()
		if !ok {
			return o, false
		}
		o = append(o, e)
	}
	return o, true
}

func (s Stack[T]) Elements() []T {
	return s.slice
}
