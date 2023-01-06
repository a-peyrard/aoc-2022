package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateQueue(t *testing.T) {
	t.Run("it should do basic queue operations", func(t *testing.T) {
		// GIVEN
		queue := CreateQueue[int]().
			Push(1).
			Push(2).
			Push(3).
			Push(4)

		// WHEN & THEN
		assert.Equal(t, 1, queue.Pop())
		assert.Equal(t, 2, queue.Pop())
		assert.Equal(t, 3, queue.Pop())
		assert.Equal(t, 4, queue.Pop())
		assert.True(t, queue.IsEmpty())
	})
}
