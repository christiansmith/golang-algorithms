package main

import "testing"

func TestGCDIterative(t *testing.T) {
	divisor := GCDIterative(6, 9)
	if divisor != 3 {
		t.Error("Expected divisor to be 3, got ", divisor)
	}
}

func BenchmarkGCDIterative(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GCDIterative(6, 9)
	}
}

func TestGCDRecursive(t *testing.T) {
	divisor := GCDRecursive(6, 9)
	if divisor != 3 {
		t.Error("Expected divisor to be 3, got", divisor)
	}
}

func BenchmarkGCDRecursive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GCDRecursive(6, 9)
	}
}
