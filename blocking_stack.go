package stack

import "sync"

type blockingStackItem[T any] struct {
	value T
	next  *blockingStackItem[T]
}

// BlockingStack implements Stack interface. Concurrent access is provided by using sync.Mutex.
type BlockingStack[T any] struct {
	mu   sync.Mutex
	head *blockingStackItem[T]
}

func NewBlockingStack[T any]() Stack[T] {
	return &BlockingStack[T]{}
}

func (bs *BlockingStack[T]) Push(val T) {
	bs.mu.Lock()

	newHead := &blockingStackItem[T]{
		value: val,
		next:  bs.head,
	}

	bs.head = newHead
	bs.mu.Unlock()
}

func (bs *BlockingStack[T]) Pop() (value T, exists bool) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	if bs.head == nil {
		return value, false
	}

	value = bs.head.value
	bs.head = bs.head.next
	return value, true
}
