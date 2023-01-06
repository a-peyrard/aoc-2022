package util

import "strconv"

func UnsafeParseInt(s string) int {
	val, _ := strconv.Atoi(s)

	return val
}
