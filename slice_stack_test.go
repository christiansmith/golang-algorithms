package main

import "testing"

func TestNewSliceStack(t *testing.T) {
	stack := NewStack()

	if len(stack) != 0 {
		t.Error(`Expected length of stack to be 0, got`, len(stack))
	}
}

func TestSliceStackPush(t *testing.T) {
	stack := NewStack()
	stack.Push(1111)
	stack.Push(2345)

	if stack[0] != 1111 {
		t.Error(`Expected the first item to be 1111, got`, stack[0])
	}
}

func TestSliceStackPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(1234)
	stack.Push(2345)
	value, _ := stack.Peek()

	if value != 2345 {
		t.Error(`Expected peeked value to be 2345, got`, value)
	}
}

func TestSliceStackPop(t *testing.T) {
	stack := NewStack()
	stack.Push(1234)
	stack.Push(2345)
	value, _ := stack.Pop()

	if value != 2345 {
		t.Error(`Expected peeked value to be 2345, got`, value)
	}
}

func BenchmarkSliceStackPush(b *testing.B) {
	b.StopTimer()
	stack := NewStack()

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		stack.Push(n)
	}
}

func BenchmarkSliceStackPeek(b *testing.B) {
	b.StopTimer()
	stack := make(Stack, b.N, b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		stack.Peek()
	}
}

func BenchmarkSliceStackPop(b *testing.B) {
	b.StopTimer()
	stack := make(Stack, b.N, b.N)
	for n := 0; n < b.N; n++ {
		stack[n] = n
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		stack.Pop()
	}
}
