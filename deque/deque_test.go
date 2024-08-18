package deque

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestNewDeque(t *testing.T) {
	t.Run("Deque type exist", func(t *testing.T) {
		var deq Deque
		assert.IsType(t, Deque{}, deq)
	})

	t.Run("New Deque", func(t *testing.T) {
		deq := NewDeque()
		assert.IsType(t, Deque{}, deq)
	})

	t.Run("New Deque empty", func(t *testing.T) {
		deq := NewDeque()
		assert.Empty(t, deq)
	})
}

func TestPushBack(t *testing.T) {
	const value = 15

	t.Run("PushBack exist", func(t *testing.T) {
		assert.NotPanics(t, func() {
			deq := NewDeque()
			deq.PushBack(value)
		})
	})

	t.Run("PushBack do something", func(t *testing.T) {
		deq := NewDeque()
		assert.Nil(t, deq.back)
		deq.PushBack(value)
		assert.NotNil(t, deq.back)
	})

	t.Run("PushBack move deque back", func(t *testing.T) {
		deq := NewDeque()
		oldBack := deq.back
		deq.PushBack(value)
		newBack := deq.back
		assert.NotEqual(t, oldBack, newBack)
	})

	t.Run("PushBack: back points to last item", func(t *testing.T) {
		deq := NewDeque()
		deq.PushBack(value)
		assert.Nil(t, (*dequeItem)(deq.back).next)
	})

	t.Run("PushBack: back points to item with correct value", func(t *testing.T) {
		deq := NewDeque()
		deq.PushBack(value)
		assert.Equal(t, value, (*dequeItem)(deq.back).value)
	})
}

func TestPopBack(t *testing.T) {
	const value = 15

	t.Run("PopBack exist", func(t *testing.T) {
		assert.NotPanics(t, func() {
			deq := NewDeque()
			_, _ = deq.PopBack()
		})
	})

	t.Run("PopBack on empty deque", func(t *testing.T) {
		deq := NewDeque()
		_, ok := deq.PopBack()
		assert.False(t, ok)
	})

	t.Run("PopBack on no empty deque", func(t *testing.T) {
		deq := NewDeque()
		deq.PushBack(value)
		_, ok := deq.PopBack()
		assert.True(t, ok)
	})

	t.Run("PopBack do something", func(t *testing.T) {
		deq := NewDeque()
		deq.PushBack(value)
		oldBack := deq.back
		deq.PopBack()
		newBack := deq.back
		assert.NotEqual(t, oldBack, newBack)
	})

	t.Run("PopBack return correct value", func(t *testing.T) {
		deq := NewDeque()
		deq.PushBack(value)
		val, _ := deq.PopBack()
		assert.Equal(t, value, val)
	})

	t.Run("PopBack: back points to last item", func(t *testing.T) {
		deq := NewDeque()
		deq.PushBack(value)
		deq.PushBack(value)
		deq.PopBack()
		assert.Nil(t, (*dequeItem)(deq.back).next)
	})
}

func TestPushBackPopBack(t *testing.T) {
	const count = 10_000

	t.Run("PushBackPopBack several times", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushBack(i)
		}
		for i := 0; i < count; i++ {
			val, ok := deq.PopBack()
			assert.True(t, ok)
			assert.Equal(t, count-1-i, val)
		}
	})

	t.Run("PushBackPopBack several times together", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushBack(i)
			val, ok := deq.PopBack()
			assert.True(t, ok)
			assert.Equal(t, i, val)
		}
	})

	t.Run("PushBackPopBack: PopBack after several times", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushBack(i)
		}
		for i := 0; i < count; i++ {
			deq.PopBack()
		}
		val, ok := deq.PopBack()
		assert.False(t, ok)
		assert.Zero(t, val)
	})

	t.Run("PushBackPopBack: PopBack after several times together", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushBack(i)
			deq.PopBack()
		}
		val, ok := deq.PopBack()
		assert.False(t, ok)
		assert.Zero(t, val)
	})
}

func TestDequeConcurrencyPushBackPopBack(t *testing.T) {
	const count = 1_000_000

	t.Run("PushBack", func(t *testing.T) {
		deq := NewDeque()

		wg := sync.WaitGroup{}
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func(value int) {
				defer wg.Done()
				deq.PushBack(value)
			}(i)
		}
		wg.Wait()

		cnt := 0
		for _, ok := deq.PopBack(); ok; _, ok = deq.PopBack() {
			cnt++
		}

		t.Log("Counter1", deq.counter1, "Counter2", deq.counter2)

		assert.Equal(t, cnt, count)
	})

	t.Run("PopBack", func(t *testing.T) {
		deq := NewDeque()

		for i := 0; i < count; i++ {
			deq.PushBack(i)
		}

		wg := sync.WaitGroup{}
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func() {
				defer wg.Done()
				deq.PopBack()
			}()
		}
		wg.Wait()

		val, ok := deq.PopBack()
		t.Log("Last val:", val, ok, "Counter1", deq.counter1, "Counter2", deq.counter2)
		assert.False(t, ok) // Queue must be empty
	})
}

func TestPushFront(t *testing.T) {
	const value = 10

	t.Run("PushFront exist", func(t *testing.T) {
		assert.NotPanics(t, func() {
			deq := NewDeque()
			deq.PushFront(value)
		})
	})

	t.Run("PushFront do something", func(t *testing.T) {
		deq := NewDeque()
		assert.Nil(t, deq.front)
		deq.PushFront(value)
		assert.NotNil(t, deq.front)
	})

	t.Run("PushFront move deque front", func(t *testing.T) {
		deq := NewDeque()
		oldFront := deq.front
		deq.PushFront(value)
		newFront := deq.front
		assert.NotEqual(t, oldFront, newFront)
	})

	t.Run("PushFront: front points to first item", func(t *testing.T) {
		deq := NewDeque()
		deq.PushFront(value)
		assert.Nil(t, (*dequeItem)(deq.front).prev)
	})

	t.Run("PushFront: front points to item with correct value", func(t *testing.T) {
		deq := NewDeque()
		deq.PushFront(value)
		assert.Equal(t, value, (*dequeItem)(deq.front).value)
	})
}

func TestPushFrontPopBack(t *testing.T) {
	const count = 100

	t.Run("PushFrontPopBack several times", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushFront(i)
		}
		for i := 0; i < count; i++ {
			val, ok := deq.PopBack()
			assert.True(t, ok)
			assert.Equal(t, i, val)
		}
	})

	t.Run("PushFrontPopBack several times together", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushFront(i)
			val, ok := deq.PopBack()
			assert.True(t, ok)
			assert.Equal(t, i, val)
		}
	})

	t.Run("PushFrontPopBack: PopBack after several times", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushFront(i)
		}
		for i := 0; i < count; i++ {
			deq.PopBack()
		}
		val, ok := deq.PopBack()
		assert.False(t, ok)
		assert.Zero(t, val)
	})

	t.Run("PushFrontPopBack: PopBack after several times together", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushFront(i)
			deq.PopBack()
		}
		val, ok := deq.PopBack()
		assert.False(t, ok)
		assert.Zero(t, val)
	})
}

func TestPopFront(t *testing.T) {
	const value = 15

	t.Run("PopFront exist", func(t *testing.T) {
		assert.NotPanics(t, func() {
			deq := NewDeque()
			_, _ = deq.PopFront()
		})
	})

	t.Run("PopFront on empty deque", func(t *testing.T) {
		deq := NewDeque()
		_, ok := deq.PopFront()
		assert.False(t, ok)
	})

	t.Run("PopFront on no empty deque", func(t *testing.T) {
		deq := NewDeque()
		deq.PushFront(value)
		_, ok := deq.PopFront()
		assert.True(t, ok)
	})

	t.Run("PopFront do something", func(t *testing.T) {
		deq := NewDeque()
		deq.PushFront(value)
		oldFront := deq.front
		deq.PopFront()
		newFront := deq.front
		assert.NotEqual(t, oldFront, newFront)
	})

	t.Run("PopFront return correct value", func(t *testing.T) {
		deq := NewDeque()
		deq.PushFront(value)
		val, _ := deq.PopFront()
		assert.Equal(t, value, val)
	})

	t.Run("PopFront: front points to first item", func(t *testing.T) {
		deq := NewDeque()
		deq.PushFront(value)
		deq.PushFront(value)
		deq.PopFront()
		assert.Nil(t, (*dequeItem)(deq.front).prev)
	})
}

func TestPushFrontPopFront(t *testing.T) {
	const count = 100

	t.Run("PushFrontPopFront several times", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushFront(i)
		}
		for i := 0; i < count; i++ {
			val, ok := deq.PopFront()
			assert.True(t, ok)
			assert.Equal(t, count-i-1, val)
		}
	})

	t.Run("PushFrontPopFront several times together", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushFront(i)
			val, ok := deq.PopFront()
			assert.True(t, ok)
			assert.Equal(t, i, val)
		}
	})

	t.Run("PushFrontPopFront: PopFront after several times", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushFront(i)
		}
		for i := 0; i < count; i++ {
			deq.PopFront()
		}
		val, ok := deq.PopFront()
		assert.False(t, ok)
		assert.Zero(t, val)
	})

	t.Run("PushFrontPopFront: PopFront after several times together", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushFront(i)
			deq.PopFront()
		}
		val, ok := deq.PopFront()
		assert.False(t, ok)
		assert.Zero(t, val)
	})
}

func TestPushBackPopFront(t *testing.T) {
	const count = 100

	t.Run("PushBackPopFront several times", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushBack(i)
		}
		for i := 0; i < count; i++ {
			val, ok := deq.PopFront()
			assert.True(t, ok)
			assert.Equal(t, i, val)
		}
	})

	t.Run("PushBackPopFront several times together", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushBack(i)
			val, ok := deq.PopFront()
			assert.True(t, ok)
			assert.Equal(t, i, val)
		}
	})

	t.Run("PushBackPopFront: PopFront after several times", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushBack(i)
		}
		for i := 0; i < count; i++ {
			deq.PopFront()
		}
		val, ok := deq.PopFront()
		assert.False(t, ok)
		assert.Zero(t, val)
	})

	t.Run("PushBackPopFront: PopFront after several times together", func(t *testing.T) {
		deq := NewDeque()
		for i := 0; i < count; i++ {
			deq.PushBack(i)
			deq.PopFront()
		}
		val, ok := deq.PopFront()
		assert.False(t, ok)
		assert.Zero(t, val)
	})
}
