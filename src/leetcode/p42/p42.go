package p42

import "math"

func trap(height []int) int {
	s := New()
	result, current := 0, 0

	for current < len(height) {
		for !s.Empty() && height[current] > height[s.Peek()] {
			top := s.Peek()
			s.Pop()
			if s.Empty() {
				break
			}
			distance := current - s.Peek() - 1
			squareHeight := int(math.Min(float64(height[current]), float64(height[s.Peek()]))) - height[top]
			result = result + distance*squareHeight
		}
		s.Push(current)
		current++
	}

	return result
}

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value int
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() int {
	if this.length == 0 {
		return 0
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() int {
	if this.length == 0 {
		return 0
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value int) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func (this *Stack) Empty() bool {
	return this.length == 0
}
