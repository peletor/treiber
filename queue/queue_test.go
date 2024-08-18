package queue

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestQueue(t *testing.T) {
	const value = 5
	const null = 0
	const count = 50

	t.Run("Push-Pop", func(t *testing.T) {
		que := NewQueue()
		que.Push(value)
		result, ok := que.Pop()
		assert.True(t, ok)
		assert.Equal(t, value, result)
	})

	t.Run("Empty Pop", func(t *testing.T) {
		que := NewQueue()
		result, ok := que.Pop()
		assert.False(t, ok)
		assert.Equal(t, null, result)
	})

	t.Run("Push Pop several times", func(t *testing.T) {
		que := NewQueue()
		for i := 0; i < count; i++ {
			que.Push(i)
		}
		for i := 0; i < count; i++ {
			result, ok := que.Pop()
			assert.True(t, ok)
			assert.Equal(t, i, result)
		}
	})

	t.Run("Pop after several Push and Pop", func(t *testing.T) {
		que := NewQueue()
		for i := 0; i < count; i++ {
			que.Push(i)
		}
		for i := 0; i < count; i++ {
			que.Pop()
		}
		result, ok := que.Pop()
		assert.False(t, ok)
		assert.Equal(t, null, result)
	})

	t.Run("Pop after several Push/Pop", func(t *testing.T) {
		que := NewQueue()
		for i := 0; i < count; i++ {
			que.Push(i)
			que.Pop()
		}
		result, ok := que.Pop()
		assert.False(t, ok)
		assert.Equal(t, null, result)
	})
}

func TestQueueConcurrency(t *testing.T) {
	const count = 50

	t.Run("Push", func(t *testing.T) {
		que := NewQueue()

		wg := sync.WaitGroup{}
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func(value int) {
				defer wg.Done()
				que.Push(value)
				que.Push(value)
				que.Push(value)
			}(i)
		}
		wg.Wait()

		cnt := 0
		for _, ok := que.Pop(); ok; _, ok = que.Pop() {
			cnt++
		}
		assert.Equal(t, cnt, count*3)
	})

	t.Run("Pop", func(t *testing.T) {
		que := NewQueue()

		for i := 0; i < count; i++ {
			que.Push(i)
		}

		wg := sync.WaitGroup{}
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func() {
				defer wg.Done()
				que.Pop()
			}()
		}
		wg.Wait()

		_, ok := que.Pop()
		assert.False(t, ok) // Queue must be empty
	})
}
