package leetcode155

// Stack represents a Stack
type Stack struct {
	vals []int
}

// NewStack allocates and returns a new stack
func NewStack() Stack {
	return Stack{vals: []int{}}
}

// Push puts a value on top of the stack
func (s *Stack) Push(val int) {
	s.vals = append(s.vals, val)
}

// Pop removes and returns the value on top of the stack
func (s *Stack) Pop() {
	s.vals = s.vals[:len(s.vals)-1]
}

// Top returns the value of the top element in the stack
func (s *Stack) Top() int {
	return s.vals[len(s.vals)-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinStack is a stack with the ability to find the value of the minimum element
// in constant time.
type MinStack struct {
	vals Stack
	mins Stack
}

// Constructor allocates a MinStack
func Constructor() MinStack {
	return MinStack{vals: NewStack(), mins: NewStack()}
}

// Push adds a value to the stack
func (ms *MinStack) Push(x int) {
	ms.vals.Push(x)
	if len(ms.mins.vals) == 0 {
		ms.mins.Push(x)
	} else {
		ms.mins.Push(min(x, ms.GetMin()))
	}
}

// Pop removes the top value from the stack
func (ms *MinStack) Pop() {
	ms.vals.Pop()
	ms.mins.Pop()
}

// Top gets the top value of the stack
func (ms *MinStack) Top() int {
	return ms.vals.vals[len(ms.vals.vals)-1]
}

// GetMin gets the minimum value in the stack
func (ms *MinStack) GetMin() int {
	return ms.mins.vals[len(ms.mins.vals)-1]
}
