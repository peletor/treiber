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

func (d *Deque) PushBack(value int) {
	newItem := unsafe.Pointer(&dequeItem{value: value, next: nil})

	for {
		back := atomic.LoadPointer(&d.back)
		(*dequeItem)(newItem).prev = back

		if back != nil {
			// Deque is not empty. dequeItem = *q.back exist
			next := atomic.LoadPointer(&(*dequeItem)(back).next)
			// if d.back is not changed in other goroutine
			if back == atomic.LoadPointer(&d.back) {
				if next == nil {
					if atomic.CompareAndSwapPointer(&(*dequeItem)(back).next, next, newItem) {
						// try to move d.back
						atomic.CompareAndSwapPointer(&d.back, back, newItem)
						return
					}
				} else {
					// try to fix d.back
					atomic.CompareAndSwapPointer(&d.back, back, next)
				}
			}
		} else {
			// Deque is empty.

			// if deque back is not changed in other goroutine
			if back == atomic.LoadPointer(&d.back) {
				if atomic.CompareAndSwapPointer(&d.back, back, newItem) {
					// try to move deque front
					atomic.CompareAndSwapPointer(&d.front, nil, newItem)
					return
				} else {
					// try to fix deque front
					atomic.CompareAndSwapPointer(&d.front, nil, newItem)
				}
			}
		}
	}
}

func (d *Deque) PopBack() (value int, ok bool) {
	for {
		back := atomic.LoadPointer(&d.back)
		if back == nil {
			// Deque is empty
			return 0, false
		}
		prev := atomic.LoadPointer(&(*dequeItem)(back).prev)

		if back == atomic.LoadPointer(&d.back) {
			// if deque has only one dequeItem
			if prev == nil {
				if atomic.CompareAndSwapPointer(&d.back, back, prev) {
					// Try to move deque front
					atomic.CompareAndSwapPointer(&d.front, back, nil)
					return (*dequeItem)(back).value, true
				}
			} else {
				// now prevItem.next == q.back
				prevItemNext := atomic.LoadPointer(&(*dequeItem)(prev).next)
				// try to make prevItem = last item (prevItem.next = nil)
				if atomic.CompareAndSwapPointer(&(*dequeItem)(prev).next, prevItemNext, nil) {
					// try to move d.back
					atomic.CompareAndSwapPointer(&d.back, back, prev)
					return (*dequeItem)(back).value, true
				} else {
					// try to fix d.back
					atomic.CompareAndSwapPointer(&d.back, back, prev)
				}

			}
		}
	}
}
