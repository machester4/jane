package main

import (
	"testing"

	"github.com/machester4/jane/lib"
)

func BenchmarkChain(b *testing.B) {
	s := "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
	lib.Initialize("chivito", "es-50")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lib.Recommend(s, "es-50", "chivito")
	}
}
