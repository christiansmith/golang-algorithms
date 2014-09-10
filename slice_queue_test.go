package main

import "testing"

func TestNewSliceQueue(t *testing.T) {
	queue := NewQueue(100000)

	if queue.head != 0 {
		t.Error(`Expected head of queue to be 0, got`, queue.head)
	}

	if queue.tail != -1 {
		t.Error(`Expected tail of queue to be -1, got`, queue.tail)
	}

	if len(queue.values) != 100000 {
		t.Error(`Expected length of queue to be 100000, got`, len(queue.values))
	}
}

func TestSliceQueuePush(t *testing.T) {
	queue := NewQueue(10)
	queue.Push(1111)

	if queue.values[0] != 1111 {
		t.Error(`Expected the first item to be 1111, got`, queue.values[0])
	}

	if queue.head != 0 {
		t.Error(`Expected head to be 0, got`, queue.head)
	}

	if queue.tail != 0 {
		t.Error(`Expected head to be 0, got`, queue.tail)
	}

	queue.Push(2222)

	if queue.values[queue.tail] != 2222 {
		t.Error(`Expected the first item to be 1111, got`, queue.values[0])
	}

	if queue.head != 0 {
		t.Error(`Expected head to be 0, got`, queue.head)
	}

	if queue.tail != 1 {
		t.Error(`Expected head to be 1, got`, queue.tail)
	}
}

func TestSliceQueuePeek(t *testing.T) {
	queue := NewQueue(10)
	queue.Push(1234)
	queue.Push(2345)
	value := queue.Peek()

	if value != 1234 {
		t.Error(`Expected peeked value to be 1234, got`, queue.values[0])
	}

	if queue.head != 0 {
		t.Error(`Expected head to be 0, got`, queue.head)
	}
}

func TestSliceQueuePop(t *testing.T) {
	queue := NewQueue(10)
	queue.Push(1234)
	queue.Push(2345)
	value := queue.Pop()

	if value != 1234 {
		t.Error(`Expected peeked value to be 2345, got`, queue.values[1])
	}

	if queue.head != 1 {
		t.Error(`Expected head to be 1, got`, queue.head)
	}
}

func BenchmarkSliceQueuePush(b *testing.B) {
	queue := NewQueue(b.N)

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		queue.Push(n)
	}
}

func BenchmarkSliceQueuePeek(b *testing.B) {
	queue := NewQueue(b.N)

	for n := 0; n < b.N; n++ {
		queue.Push(n)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		queue.Peek()
	}
}

func BenchmarkSliceQueuePop(b *testing.B) {
	queue := NewQueue(b.N)

	for n := 0; n < b.N; n++ {
		queue.Push(n)
	}

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		queue.Pop()
	}
}
