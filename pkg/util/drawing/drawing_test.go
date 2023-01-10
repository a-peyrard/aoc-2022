package drawing

import (
	"aoc2022/pkg/util/collection"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDrawing(t *testing.T) {
	t.Run("it should create a nice drawing :)", func(t *testing.T) {
		// GIVEN
		drawing := InitDrawing(3, 4)

		// WHEN
		/*
		   1..
		   .2.
		   ..0
		   sX.
		*/
		res := drawing.
			Fill('.').
			DrawAt('s', 0, 0).
			DrawAt('X', 1, 0).
			DrawAt('1', 0, 3).
			DrawAt('2', 1, 2).
			DrawAt('0', 2, 1).
			String()

		// THEN
		assert.Equal(t, `1..
.2.
..0
sX.`, res)
	})

	t.Run("it should create a nice drawing from top to bottom", func(t *testing.T) {
		// GIVEN
		drawing := InitDrawingTopToBottom(3, 4)

		// WHEN
		/*
		   1..
		   .2.
		   ..0
		   sX.
		*/
		res := drawing.
			Fill('.').
			DrawAt('s', 0, 0).
			DrawAt('X', 1, 0).
			DrawAt('1', 0, 3).
			DrawAt('2', 1, 2).
			DrawAt('0', 2, 1).
			String()

		// THEN
		assert.Equal(t, `sX.
..0
.2.
1..`, res)
	})

	t.Run("it should create a nice drawing with a decentralized center", func(t *testing.T) {
		// GIVEN
		drawing := InitDrawingWithCenter(3, 4, 1, 1)

		// WHEN
		/*
		   1..
		   .2.
		   .s0
		   .X.
		*/
		res := drawing.
			Fill('.').
			DrawAt('s', 0, 0).
			DrawAt('X', 0, -1).
			DrawAt('1', -1, 2).
			DrawAt('2', 0, 1).
			DrawAt('0', 1, 0).
			String()

		// THEN
		assert.Equal(t, `1..
.2.
.s0
.X.`, res)
	})

	t.Run("it should retrieve value of a 'cell'", func(t *testing.T) {
		// GIVEN
		/*
		   1..
		   .2.
		   ..0
		   sX.
		*/
		drawing := InitDrawing(3, 4).
			Fill('.').
			DrawAt('s', 0, 0).
			DrawAt('X', 1, 0).
			DrawAt('1', 0, 3).
			DrawAt('2', 1, 2).
			DrawAt('0', 2, 1)

		// WHEN
		val := drawing.ValueAt(0, 3)

		// THEN
		assert.Equal(t, byte('1'), val)
	})
}

func TestDrawMatrix(t *testing.T) {
	t.Run("it should draw a matrix", func(t *testing.T) {
		// GIVEN
		matrix := collection.CreateMatrix(3, 2, 0)
		/*
			123
			456
		*/
		matrix.PutAt(0, 0, 1)
		matrix.PutAt(1, 0, 2)
		matrix.PutAt(2, 0, 3)
		matrix.PutAt(0, 1, 4)
		matrix.PutAt(1, 1, 5)
		matrix.PutAt(2, 1, 6)

		// WHEN
		drawing := DrawMatrix(&matrix, func(val int, x int, y int) byte {
			return byte('0' + val)
		}).String()

		assert.Equal(t, `123
456`, drawing)
	})
}
