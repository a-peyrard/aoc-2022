package util

import "golang.org/x/exp/constraints"

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T constraints.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
