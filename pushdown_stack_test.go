package main

import "testing"

func TestPushdownStack(t *testing.T) {
	stack := NewPushdownStack()
	data := []int{7, 5, 6, 4, 8, 3}

	for _, v := range data {
		stack.Push(v)
	}

	for i := len(data) - 1; stack.IsEmpty() != true; i-- {
		result, expected := stack.Pop(), data[i]
		if result != expected {
			t.Error("Expected popped value to be", expected, "got", result)
		}
	}
}

var pstack = NewPushdownStack()

func BenchmarkPushdownStackPush(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pstack.Push(n)
	}
}

func BenchmarkPushdownStackPop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		pstack.Push(n)
	}
}
