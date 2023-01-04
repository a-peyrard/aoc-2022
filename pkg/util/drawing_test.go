package util

import (
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
}
