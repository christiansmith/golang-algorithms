package main

type CappedStack struct {
	values []int
	top    int
}

func NewCappedStack(n int) *CappedStack {
	return &CappedStack{make([]int, n), -1}
}

func (stack *CappedStack) Push(value int) bool {
	if stack.top+1 == cap(stack.values) {
		return false
	}

	stack.top++
	stack.values[stack.top] = value
	return true
}

func (stack *CappedStack) Peek() (int, bool) {
	if len(stack.values) == 0 {
		return 0, false
	}
	return stack.values[stack.top], true
}

func (stack *CappedStack) Pop() (int, bool) {
	if len(stack.values) == 0 {
		return 0, false
	}

	value := stack.values[stack.top]
	stack.top--
	return value, true
}
