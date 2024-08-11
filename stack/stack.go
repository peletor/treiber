package stack

import (
	"sync/atomic"
	"unsafe"
)

type stackItem struct {
	value int
	next  unsafe.Pointer
}
type Stack struct {
	head unsafe.Pointer
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(value int) {
	newNode := &stackItem{value: value}

	for {
		head := atomic.LoadPointer(&s.head)
		newNode.next = head

		if atomic.CompareAndSwapPointer(&s.head, head, unsafe.Pointer(newNode)) {
			return
		}
	}
}

func (s *Stack) Pop() (value int, Ok bool) {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return 0, false
		}

		next := atomic.LoadPointer(&(*stackItem)(head).next)
		if atomic.CompareAndSwapPointer(&s.head, head, next) {
			return (*stackItem)(head).value, true
		}
	}
}

func (s *Stack) Top() (value int, Ok bool) {
	for {
		head := atomic.LoadPointer(&s.head)
		if head == nil {
			return 0, false
		}

		// Try to swap head with itself
		if atomic.CompareAndSwapPointer(&s.head, head, head) {
			return (*stackItem)(head).value, true
		}
	}
}
