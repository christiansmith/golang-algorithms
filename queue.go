package main

// Node is defined already

type LinkedListQueue struct {
	Head, Tail *Node
}

func NewLinkedListQueue() *LinkedListQueue {
	queue := &LinkedListQueue{
		Head: &Node{Key: 0},
	}

	queue.Tail = queue.Head
	queue.Head.Next = queue.Tail

	return queue
}

func (queue *LinkedListQueue) Insert(v int) {
	queue.Tail.Next = &Node{Key: v}
	queue.Tail = queue.Tail.Next
	queue.Tail.Next = queue.Tail
}

func (queue *LinkedListQueue) Remove() int {
	node := queue.Head.Next
	queue.Head.Next = node.Next
	return node.Key
}

func (queue *LinkedListQueue) IsEmpty() bool {
	return queue.Head.Next == queue.Tail
}
