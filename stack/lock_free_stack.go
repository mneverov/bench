package stack

import (
	"sync/atomic"
)

// LockFreeStack implements Stack interface using Treiber's algorithm https://en.wikipedia.org/wiki/Treiber_stack.
type LockFreeStack[T any] struct {
	head atomic.Pointer[lockFreeStackItem[T]]
}

type lockFreeStackItem[T any] struct {
	value T
	next  *lockFreeStackItem[T]
}

func NewLockFreeStack[T any]() Stack[T] {
	return &LockFreeStack[T]{}
}

func (s *LockFreeStack[T]) Push(value T) {
	node := &lockFreeStackItem[T]{value: value}

	for {
		head := s.head.Load()
		node.next = head

		if s.head.CompareAndSwap(head, node) {
			return
		}
	}
}

func (s *LockFreeStack[T]) Pop() (value T, exists bool) {
	for {
		head := s.head.Load()
		if head == nil {
			return value, false
		}

		next := head.next
		if s.head.CompareAndSwap(head, next) {
			return head.value, true
		}
	}
}
