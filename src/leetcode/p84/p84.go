package p84

func largestRectangleArea(heights []int) int {
	return 0
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
