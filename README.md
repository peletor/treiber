# Treiber
Go's concurrent and lock-free data structures are based on a concept known as a Treiber stack.

This package contains data structures:
- Stack - [Treiber stack](https://en.wikipedia.org/wiki/Treiber_stack)
- Queue - [Michael-Scott Queue](https://www.cs.rochester.edu/~scott/papers/1996_PODC_queues.pdf)
- Deque - double-ended queue

## Stack
implements methods:

- Push – adds an element to the collection.
- Pop – removes the most recently added element.
- Top – retrieves the value from the top of the stack.

## Queue
implements methods:
- Push – adds an element to the end of the queue.
- Pop – removes an element from the beginning of the queue.

## Deque
implements methods:
- PushBack – adds an element to the end of the deque.
- PushFront – adds an element to the beginning of the deque.
- PopBack – removes an element from the end of the deque.
- PopFront – removes an element from the beginning of the deque.