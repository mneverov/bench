package stack

type Stack[T any] interface {
	Push(v T)
	Pop() (T, bool)
}

type item[T any] struct {
	value T
	next  *item[T]
}

type RegularStack[T any] struct {
	head *item[T]
}

func NewRegularStack[T any]() Stack[T] {
	return &RegularStack[T]{}
}

func (s *RegularStack[T]) Push(value T) {
	s.head = &item[T]{value: value, next: s.head}
}

func (s *RegularStack[T]) Pop() (value T, exists bool) {
	if s.head == nil {
		return value, false
	}

	value = s.head.value
	s.head = s.head.next
	return value, true
}
