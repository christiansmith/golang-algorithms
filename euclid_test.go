package main

import "testing"

func TestGCD(t *testing.T) {
	divisor := GCD(6, 9)
	if divisor != 3 {
		t.Error("Expected divisor to be 3, got ", divisor)
	}
}

func BenchmarkGCD(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GCD(6, 9)
	}
}
