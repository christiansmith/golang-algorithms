package main

import "testing"

func TestLinkedListQueue(t *testing.T) {
	queue := NewLinkedListQueue()
	data := []int{7, 5, 6, 4, 8, 3}

	for _, v := range data {
		queue.Insert(v)
	}

	for i := 0; queue.IsEmpty() != true; i++ {
		result, expected := queue.Remove(), data[i]
		if result != expected {
			t.Error("Expected removed value to be", expected, "got", result)
		}
	}
}

var queue = NewLinkedListQueue()

func BenchmarkLinkedListQueueInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		queue.Insert(n)
	}
}

func BenchmarkLinkedListQueueRemove(b *testing.B) {
	for n := 0; n < b.N; n++ {
		queue.Remove()
	}
}
