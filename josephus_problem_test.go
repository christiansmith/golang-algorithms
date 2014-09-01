package main

import "testing"

func TestJosephus(t *testing.T) {
	last := Josephus(9, 5)
	if last.Key != 8 {
		t.Error("Expected 8, got", last.Key)
	}
}

func BenchmarkJosephus(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Josephus(9, 5)
	}
}
