package queue

import (
	"sync/atomic"
	"unsafe"
)

type queueItem struct {
	value int
	next  unsafe.Pointer
}

type Queue struct {
	head unsafe.Pointer
	tail unsafe.Pointer
}

func NewQueue() Queue {
	firstItem := &queueItem{}
	return Queue{
		head: unsafe.Pointer(firstItem),
		tail: unsafe.Pointer(firstItem),
	}
}

func (q *Queue) Push(value int) {
	newItem := &queueItem{value: value}

	for {
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*queueItem)(tail).next)

		// if queue tail is not changed in other goroutine
		if tail == atomic.LoadPointer(&q.tail) {
			if next == nil {
				if atomic.CompareAndSwapPointer(&(*queueItem)(tail).next, next, unsafe.Pointer(newItem)) {
					// try to move queue tail
					atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(newItem))
					return
				} else {
					// try to fix queue tail
					atomic.CompareAndSwapPointer(&q.tail, tail, unsafe.Pointer(newItem))
				}

			}
		}
	}
}

func (q *Queue) Pop() (value int, ok bool) {
	for {
		head := atomic.LoadPointer(&q.head)
		tail := atomic.LoadPointer(&q.tail)
		next := atomic.LoadPointer(&(*queueItem)(head).next)

		// if queue head is not changed in other goroutine
		if head == atomic.LoadPointer(&q.head) {
			if head == tail {
				if next == nil {
					// queue is empty
					return 0, false
				} else {
					// fix queue tail
					atomic.CompareAndSwapPointer(&q.tail, tail, next)
				}
			} else {
				value := (*queueItem)(next).value
				if atomic.CompareAndSwapPointer(&q.head, head, next) {
					// head has been changed successfully
					return value, true
				}
			}
		}
	}
}
