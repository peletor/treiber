package deque

import (
	"github.com/stretchr/testify/assert"
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
	const count = 100

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
