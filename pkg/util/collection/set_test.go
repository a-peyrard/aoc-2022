package collection

import (
	"aoc2022/pkg/util/geo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("it should work as expected", func(t *testing.T) {
		// GIVEN
		set := NewSet[int]()

		// WHEN & THEN

		// Add some values to the set.
		set.Add(3)
		set.Add(42)
		set.Add(5)

		// Check if a value is in the set.
		assert.True(t, set.Contains(3))
		assert.True(t, set.Contains(42))
		assert.True(t, set.Contains(5))
		assert.False(t, set.Contains(1))

		// Remove a value from the set.
		set.Remove(42)

		// Print the values in the set.
		s := set.Values()
		assert.Contains(t, s, 3)
		assert.Contains(t, s, 5)
	})

	t.Run("it should work with coordinate", func(t *testing.T) {
		// GIVEN
		set := NewSet[geo.Coordinate]()

		// WHEN & THEN

		// Add some values to the set.
		set.Add(geo.Coordinate{X: 3, Y: 4})
		set.Add(geo.Coordinate{X: 5, Y: 7})

		// Check if a value is in the set.
		assert.True(t, set.Contains(geo.Coordinate{X: 3, Y: 4}))
	})
}
