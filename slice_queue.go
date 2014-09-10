package main

type Queue struct {
	values []int
	head   int
	tail   int
}

func NewQueue(n int) *Queue {
	return &Queue{make([]int, n), 0, -1}
}

func (queue *Queue) Push(value int) int {
	queue.tail++
	queue.values[queue.tail] = value
	return queue.tail - 1
}

func (queue *Queue) Peek() int {
	if queue.head <= queue.tail {
		return queue.values[queue.head]
	}
	return 0
}

func (queue *Queue) Pop() int {
	if queue.head <= queue.tail {
		value := queue.values[queue.head]
		queue.head++
		return value
	}
	return 0
}

func (queue *Queue) Length() int {
	switch {
	case queue.head > queue.tail:
		return 0
	case queue.head == queue.tail:
		return 1
	default:
		return (queue.tail - queue.head) + 1
	}
}

func (queue *Queue) IsEmpty() bool {
	return queue.head > queue.tail
}
