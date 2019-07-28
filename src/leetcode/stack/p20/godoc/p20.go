package godoc

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

func isValid(s string) bool {
	array := []rune(s)
	if len(array)%2 != 0 {
		return false
	}
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack Stack
	for index := 0; index < len(array); index++ {
		r := array[index]
		_, ok := mapping[r]
		if ok {
			pop := stack.Pop()
			l, ok := pop.(rune)
			if ok {
				if mapping[r] != l {
					return false
				}
			} else {
				return false
			}
		} else {
			stack.Push(r)
		}
	}
	if stack.Len() == 0 {
		return true
	} else {
		return false
	}

}
