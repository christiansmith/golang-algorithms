package main

// We don't need struct to wrap the slice,
// since there are no properties to keep
// track of. Instead, simply alias the type
// "slice of int".
type Stack []int

// This is unnecessary, as we can simply
// declare of variable of type "Stack".
// Also, we can use `make(Stack, len, cap)`
// to preallocate a sufficiently sized array
// if we know up front approximately how
// large we need the underlying array to be.
func NewStack() Stack {
	var s Stack
	return s
}

// Push can rely on the built in `append` function.
func (s *Stack) Push(value int) {
	*s = append(*s, value)
}

// Because a zero value can be a valid item on the
// stack, Peek and Pop return two values. The first
// is the value of the last item, and the second
// indicates whether or not the array is not empty.

func (s *Stack) Peek() (int, bool) {
	// dereference the pointer
	stack := *s

	// ensure the stack is not empty
	if len(stack) == 0 {
		return 0, false
	}

	// peek at the value
	return stack[len(stack)-1], true
}

func (s *Stack) Pop() (int, bool) {
	// dereference the pointer
	stack := *s

	// ensure the stack is not empty
	if len(stack) == 0 {
		return 0, false
	}

	// pop the value
	index := len(stack) - 1
	value := stack[index]
	stack = stack[:index]
	return value, true
}

func (s *Stack) Length() int {
	return len(*s)
}

func (s *Stack) IsEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}
