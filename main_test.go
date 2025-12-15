package main

import (
	"testing"
)

func BenchmarkAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		main()
	}
}
