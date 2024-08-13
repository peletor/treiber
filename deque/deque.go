package deque

import (
	"sync/atomic"
	"unsafe"
)

type dequeItem struct {
	value int
	prev  unsafe.Pointer
	next  unsafe.Pointer
}

type Deque struct {
	front unsafe.Pointer
	back  unsafe.Pointer
}

func NewDeque() Deque {
	return Deque{}
}

func (q *Deque) PushBack(value int) {
	newItem := &dequeItem{value: value, next: nil}

	for {
		back := atomic.LoadPointer(&q.back)
		newItem.prev = back

		if atomic.CompareAndSwapPointer(&q.back, back, unsafe.Pointer(newItem)) {
			return
		}
	}
}

func (q *Deque) PopBack() (value int, ok bool) {
	for {
		back := atomic.LoadPointer(&q.back)
		if back == nil {
			return 0, false
		}
		prev := atomic.LoadPointer(&(*dequeItem)(back).prev)

		if atomic.CompareAndSwapPointer(&q.back, back, prev) {
			return (*dequeItem)(back).value, true
		}
	}
}
