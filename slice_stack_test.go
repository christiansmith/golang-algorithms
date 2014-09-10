package main

import "testing"

func TestNewSliceStack(t *testing.T) {
	stack := NewStack(100000)

	if stack.top != -1 {
		t.Error(`Expected top of stack to be -1, got`, stack.top)
	}

	if len(stack.values) != 100000 {
		t.Error(`Expected length of stack to be 100000, got`, len(stack.values))
	}
}

func TestSliceStackPush(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1111)
	stack.Push(2345)

	if stack.values[0] != 1111 {
		t.Error(`Expected the first item to be 1111, got`, stack.values[0])
	}

	if stack.top != 1 {
		t.Error(`Expected top to be 1, got`, stack.top)
	}
}

func TestSliceStackPeek(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1234)
	stack.Push(2345)
	value := stack.Peek()

	if value != 2345 {
		t.Error(`Expected peeked value to be 2345, got`, stack.values[1])
	}

	if stack.top != 1 {
		t.Error(`Expected top to be 1, got`, stack.top)
	}
}

func TestSliceStackPop(t *testing.T) {
	stack := NewStack(10)
	stack.Push(1234)
	stack.Push(2345)
	value := stack.Pop()

	if value != 2345 {
		t.Error(`Expected peeked value to be 2345, got`, stack.values[1])
	}

	if stack.top != 0 {
		t.Error(`Expected top to be 0, got`, stack.top)
	}
}

func BenchmarkSliceStackPush(b *testing.B) {
	stack := NewStack(b.N)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		stack.Push(n)
	}
}

func BenchmarkSliceStackPeek(b *testing.B) {
	stack := NewStack(b.N)

	for n := 0; n < b.N; n++ {
		stack.Push(n)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		stack.Peek()
	}
}

func BenchmarkSliceStackPop(b *testing.B) {
	stack := NewStack(b.N)

	for n := 0; n < b.N; n++ {
		stack.Push(n)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		stack.Pop()
	}
}
