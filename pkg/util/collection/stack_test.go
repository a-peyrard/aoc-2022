package collection

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Run("it should have a working length", func(t *testing.T) {
		// GIVEN
		stack := CreateStack[int]()
		stack.Push(10).Push(20).Push(30)

		// WHEN
		got := stack.Length()

		// THEN
		want := 3
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Length() = %v, want %v", got, want)
		}
	})

	t.Run("it should pop last insertion", func(t *testing.T) {
		// GIVEN
		stack := CreateStack[int]()
		stack.Push(10).Push(20).Push(30)

		// WHEN
		got := stack.Pop()

		// THEN
		want := 30
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Pop() = %v, want %v", got, want)
		}
	})

	t.Run("it should remove till predicate is false", func(t *testing.T) {
		// GIVEN
		type pair struct {
			value int
			index int
		}
		stack := CreateStack[*pair]()
		stack.Push(&pair{10, 0})
		stack.Push(&pair{8, 0})
		stack.Push(&pair{4, 0})
		stack.Push(&pair{3, 0})
		stack.Push(&pair{1, 0})

		// WHEN
		stack.RemoveTill(func(p *pair) bool {
			return p.value < 5
		})

		// THEN
		assert.Equal(t, 2, stack.Length())
		assert.Equal(t, 8, stack.Peek().value)
	})

	t.Run("it should remove till predicate is false and return if stack is empty", func(t *testing.T) {
		// GIVEN
		type pair struct {
			value int
			index int
		}
		stack := CreateStack[*pair]()
		stack.Push(&pair{10, 0})
		stack.Push(&pair{8, 0})
		stack.Push(&pair{4, 0})
		stack.Push(&pair{3, 0})
		stack.Push(&pair{1, 0})

		// WHEN
		res := stack.RemoveTill(func(p *pair) bool {
			return p.value < 99
		})

		// THEN
		assert.Equal(t, true, res)
		assert.Equal(t, 0, stack.Length())
	})

	t.Run("it should clear the stack", func(t *testing.T) {
		// GIVEN
		stack := CreateStack[int]()
		stack.Push(10).Push(20).Push(30)

		// WHEN
		stack.Clear()

		// THEN
		assert.Equal(t, true, stack.IsEmpty())
	})

	t.Run("it should reverse a stack", func(t *testing.T) {
		// GIVEN
		stack := CreateStack[int]()
		stack.Push(10).Push(20).Push(30)

		// WHEN
		stack.Reverse()

		// THEN
		got1 := stack.Pop()
		got2 := stack.Pop()
		got3 := stack.Pop()
		want1 := 10
		want2 := 20
		want3 := 30
		if !reflect.DeepEqual(got1, want1) {
			t.Errorf("Pop1() = %v, want %v", got1, want1)
		}
		if !reflect.DeepEqual(got2, want2) {
			t.Errorf("Pop2() = %v, want %v", got2, want2)
		}
		if !reflect.DeepEqual(got3, want3) {
			t.Errorf("Pop3() = %v, want %v", got3, want3)
		}
	})
}
