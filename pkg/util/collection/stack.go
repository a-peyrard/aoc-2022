package collection

/*
	Try to implement a basic stack data structure.

	Also found this interesting implementation on SO:
type Stack[T any] struct {
	Push   func(T)
	Pop    func() T
	Length func() int
}

func NewStack[T any]() Stack[T] {
	slice := make([]T, 0)
	return Stack[T]{
		Push: func(i T) {
			slice = append(slice, i)
		},
		Pop: func() T {
			res := slice[len(slice)-1]
			slice = slice[:len(slice)-1]
			return res
		},
		Length: func() int {
			return len(slice)
		},
	}
}
https://stackoverflow.com/questions/28541609/looking-for-reasonable-stack-implementation-in-golang

*/

type Stack[T any] struct {
	inner []T
}

func CreateStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func (s *Stack[T]) Push(element T) *Stack[T] {
	s.inner = append(s.inner, element)
	return s
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		panic("Stack is empty!")
	}
	head := s.inner[len(s.inner)-1]
	s.inner = s.inner[:len(s.inner)-1]
	return head
}

func (s *Stack[T]) Peek() T {
	return s.inner[len(s.inner)-1]
}

func (s *Stack[T]) Length() int {
	return len(s.inner)
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack[T]) Reverse() *Stack[T] {
	length := len(s.inner)
	for i := 0; i < length/2; i++ {
		s.inner[i], s.inner[length-1-i] = s.inner[length-1-i], s.inner[i]
	}
	return s
}
