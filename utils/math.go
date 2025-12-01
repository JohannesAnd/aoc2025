package utils

import "golang.org/x/exp/constraints"

func Mod[T constraints.Integer](n, m T) T {
	if m == 0 {
		panic("modulo by zero")
	}

	r := n % m

	if r < 0 {
		r += m
	}

	return r
}
