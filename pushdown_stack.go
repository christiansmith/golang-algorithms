package main

// Node is defined already

type PushdownStack struct {
	Head, Z *Node
}

func NewPushdownStack() *PushdownStack {
	stack := &PushdownStack{
		Head: &Node{},
		Z:    &Node{},
	}

	stack.Head.Next = stack.Z
	stack.Head.Key = 0
	stack.Z.Next = stack.Z

	return stack
}

func (stack *PushdownStack) Push(v int) {
	t := &Node{v, stack.Head.Next}
	stack.Head.Next = t
}

func (stack *PushdownStack) Pop() int {
	t := stack.Head.Next
	stack.Head.Next = t.Next
	return t.Key
}

func (stack *PushdownStack) IsEmpty() bool {
	return stack.Head.Next == stack.Z
}
