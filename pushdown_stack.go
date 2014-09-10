package main

type StackItem struct {
	Value int
	Next  *StackItem
}

type PushdownStack struct {
	Top *StackItem
}

func NewPushdownStack() *PushdownStack {
	return &PushdownStack{}
}

func (stack *PushdownStack) Push(value int) {
	stack.Top = &StackItem{value, stack.Top}
}

func (stack *PushdownStack) Peek() int {
	return stack.Top.Value
}

func (stack *PushdownStack) Pop() int {
	top := stack.Top

	if top != nil {
		stack.Top = top.Next
		return top.Value
	}

	return 0
}

func (stack *PushdownStack) IsEmpty() bool {
	return stack.Top == nil
}
