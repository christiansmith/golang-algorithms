package main

import "testing"

func TestNewCappedStack(t *testing.T) {
	stack := NewCappedStack(100000)

	if stack.top != -1 {
		t.Error(`Expected top of stack to be -1, got`, stack.top)
	}

	if len(stack.values) != 100000 {
		t.Error(`Expected length of stack to be 100000, got`, len(stack.values))
	}
}

func TestCappedStackPush(t *testing.T) {
	stack := NewCappedStack(10)
	stack.Push(1111)
	stack.Push(2345)

	if stack.values[0] != 1111 {
		t.Error(`Expected the first item to be 1111, got`, stack.values[0])
	}

	if stack.top != 1 {
		t.Error(`Expected top to be 1, got`, stack.top)
	}
}

func TestCappedStackPeek(t *testing.T) {
	stack := NewCappedStack(10)
	stack.Push(1234)
	stack.Push(2345)
	value, _ := stack.Peek()

	if value != 2345 {
		t.Error(`Expected peeked value to be 2345, got`, stack.values[1])
	}

	if stack.top != 1 {
		t.Error(`Expected top to be 1, got`, stack.top)
	}
}

func TestCappedStackPop(t *testing.T) {
	stack := NewCappedStack(10)
	stack.Push(1234)
	stack.Push(2345)
	value, _ := stack.Pop()

	if value != 2345 {
		t.Error(`Expected peeked value to be 2345, got`, stack.values[1])
	}

	if stack.top != 0 {
		t.Error(`Expected top to be 0, got`, stack.top)
	}
}

func BenchmarkCappedStackPush(b *testing.B) {
	b.StopTimer()
	stack := NewCappedStack(b.N)

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		stack.Push(n)
	}
}

func BenchmarkCappedStackPeek(b *testing.B) {
	b.StopTimer()
	stack := NewCappedStack(b.N)
	for n := 0; n < b.N; n++ {
		stack.Push(n)
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		stack.Peek()
	}
}

func BenchmarkCappedStackPop(b *testing.B) {
	b.StopTimer()
	stack := NewCappedStack(b.N)
	for n := 0; n < b.N; n++ {
		stack.Push(n)
	}

	b.StartTimer()
	for n := 0; n < b.N; n++ {
		stack.Pop()
	}
}
