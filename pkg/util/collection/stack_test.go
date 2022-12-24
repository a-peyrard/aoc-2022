package collection

import (
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
