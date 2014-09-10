package main

type Stack struct {
	values []int
	top    int
}

func NewStack(n int) *Stack {
	return &Stack{make([]int, n), -1}
}

func (stack *Stack) Push(value int) int {
	stack.top++
	stack.values[stack.top] = value
	return stack.top
}

func (stack *Stack) Peek() int {
	if stack.top >= 0 {
		return stack.values[stack.top]
	}
	return 0
}

func (stack *Stack) Pop() int {
	if stack.top >= 0 {
		value := stack.values[stack.top]
		stack.top--
		return value
	}
	return 0
}
