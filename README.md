# Treiber
Go's concurrent and scalable lock-free data structures are based on the concept of a Treiber stack.

This package contains data structures:
- Stack - [Treiber stack](https://en.wikipedia.org/wiki/Treiber_stack)
- Queue - [Michael-Scott Queue](https://www.cs.rochester.edu/~scott/papers/1996_PODC_queues.pdf)

## Stack
implements methods:

- Push – adds an element to the collection.
- Pop – removes the most recently added element.
- Top – retrieves the value from the top of the stack.

## Queue
implements methods:
- Push – adds an element to the end of the queue.
- Pop – removes an element from the beginning of the queue.

