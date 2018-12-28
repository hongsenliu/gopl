package main

import (
	"testing"
)

func BenchmarkM1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m1()
	}
}

func BenchmarkM2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		m2()
	}
}
