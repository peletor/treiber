package stack

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestStack(t *testing.T) {
	const value = 5
	const null = 0

	t.Run("Push-Top", func(t *testing.T) {
		st := NewStack()
		st.Push(value)
		result, ok := st.Top()
		assert.True(t, ok)
		assert.Equal(t, value, result)
	})

	t.Run("Push-Pop", func(t *testing.T) {
		st := NewStack()
		st.Push(value)
		result, ok := st.Pop()
		assert.True(t, ok)
		assert.Equal(t, value, result)
	})

	t.Run("Empty Top", func(t *testing.T) {
		st := NewStack()
		result, ok := st.Top()
		assert.False(t, ok)
		assert.Equal(t, null, result)
	})

	t.Run("Empty Pop", func(t *testing.T) {
		st := NewStack()
		result, ok := st.Pop()
		assert.False(t, ok)
		assert.Equal(t, null, result)
	})

	const count = 10
	t.Run("Push Pop several times", func(t *testing.T) {
		st := NewStack()
		for i := 0; i < count; i++ {
			st.Push(i)
		}
		for i := 0; i < count; i++ {
			result, ok := st.Pop()
			assert.True(t, ok)
			assert.Equal(t, count-i-1, result)
		}
	})

	t.Run("Top after Push Pop several times", func(t *testing.T) {
		st := NewStack()
		for i := 0; i < count; i++ {
			st.Push(i)
		}
		for i := 0; i < count; i++ {
			st.Pop()
		}
		result, ok := st.Top()
		assert.False(t, ok)
		assert.Equal(t, null, result)
	})

	t.Run("Pop after Push Pop several times", func(t *testing.T) {
		st := NewStack()
		for i := 0; i < count; i++ {
			st.Push(i)
		}
		for i := 0; i < count; i++ {
			st.Pop()
		}
		result, ok := st.Pop()
		assert.False(t, ok)
		assert.Equal(t, null, result)
	})
}

func TestStackConcurrency(t *testing.T) {
	const count = 50

	t.Run("Push", func(t *testing.T) {
		st := NewStack()

		wg := sync.WaitGroup{}
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func(value int) {
				defer wg.Done()
				st.Push(value)
				st.Push(value)
				st.Push(value)
			}(i)
		}
		wg.Wait()

		cnt := 0
		for _, ok := st.Pop(); ok; _, ok = st.Pop() {
			cnt++
		}
		assert.Equal(t, cnt, count*3)
	})

	t.Run("Pop", func(t *testing.T) {
		st := NewStack()

		for i := 0; i < count; i++ {
			st.Push(i)
		}

		wg := sync.WaitGroup{}
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func() {
				defer wg.Done()
				st.Pop()
			}()
		}
		wg.Wait()

		_, ok := st.Pop()
		assert.False(t, ok) // Stack must be empty
	})
}
