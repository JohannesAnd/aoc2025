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

func IsEven(n int) bool {
	return n%2 == 0
}

func IsOdd(n int) bool {
	return n%2 == 1
}
