package p9

type CQueue struct {
	stack1 stack
	stack2 stack
}

func Constructor() CQueue {
	return CQueue{stack1: stack{}, stack2: stack{}}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1 = this.stack1.push(value)
}

func (this *CQueue) DeleteHead() int {
	var pop = -1
	if this.stack2.empty() {
		for !this.stack1.empty() {
			this.stack1, pop = this.stack1.pop()
			this.stack2 = this.stack2.push(pop)
		}
	}
	if !this.stack2.empty() {
		this.stack2, pop = this.stack2.pop()
	}
	return pop
}

type stack []int

func (s stack) len() int {
	return len(s)
}
func (s stack) empty() bool {
	return s.len() == 0
}

func (s stack) push(v int) stack {
	return append(s, v)
}

func (s stack) pop() (stack, int) {
	return s[:s.len()-1], s[s.len()-1]
}
