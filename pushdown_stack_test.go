package main

import "testing"

func TestPushdownStackPush(t *testing.T) {
	stack := NewPushdownStack()
	stack.Push(1324)
	stack.Push(2435)

	if stack.Top.Value != 2435 {
		t.Error(`Expected the top of the stack to be 2435, got`, stack.Top.Value)
	}

	if stack.Top.Next.Value != 1324 {
		t.Error(`Expected next in stack to be 1324, got,`, stack.Top.Next.Value)
	}
}

func TestPushdownStackPeek(t *testing.T) {
	stack := NewPushdownStack()
	stack.Push(1324)
	stack.Push(2435)
	value := stack.Peek()

	if value != 2435 {
		t.Error(`Expected to peek at 2435, got`, stack.Top.Value)
	}
}

func TestPushdownStackPop(t *testing.T) {
	stack := NewPushdownStack()
	stack.Push(1324)
	stack.Push(2435)
	value := stack.Pop()

	if value != 2435 {
		t.Error(`Expected to peek at 2435, got`, stack.Top.Value)
	}

	if stack.Top.Value != 1324 {
		t.Error(`Expected top of stack to be 1324, got`, stack.Top.Value)
	}
}

func BenchmarkPushdownStackPush(b *testing.B) {
	stack := NewPushdownStack()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		stack.Push(n)
	}
}

func BenchmarkPushdownStackPeek(b *testing.B) {
	stack := NewPushdownStack()
	stack.Push(1324)
	stack.Push(2435)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		stack.Peek()
	}
}

func BenchmarkPushdownStackPop(b *testing.B) {
	stack := NewPushdownStack()

	for n := 0; n < b.N; n++ {
		stack.Push(n)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		stack.Pop()
	}
}
