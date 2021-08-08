package p232

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
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
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

type MyQueue struct {
	stackA *Stack
	stackB *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{stackA: New(), stackB: New()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	if this.stackA.Len() != 0 {
		this.stackA.Push(x)
	} else if this.stackB.Len() != 0 {
		for this.stackB.Len() > 0 {
			this.stackA.Push(this.stackB.Pop())
		}
		this.stackA.Push(x)
	} else {
		this.stackA.Push(x)
	}
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.stackA.Len() != 0 {
		for this.stackA.Len() > 0 {
			this.stackB.Push(this.stackA.Pop())
		}
		return this.stackB.Pop().(int)
	} else {
		return this.stackB.Pop().(int)
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.stackA.Len() != 0 {
		node := this.stackA.top
		for node.prev != nil {
			node = node.prev
		}
		return node.value.(int)
	} else {
		return this.stackB.top.value.(int)
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.stackA.Len() == 0 && this.stackB.Len() == 0 {
		return true
	}
	return false
}
