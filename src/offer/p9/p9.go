package p9

type CQueue1 struct {
	stack1 stack
	stack2 stack
}

func Constructor1() CQueue {
	return CQueue{stack1: stack{}, stack2: stack{}}
}

func (this *CQueue) AppendTail1(value int) {
	if !this.stack2.empty() {
		var pop int
		for !this.stack2.empty() {
			this.stack2, pop = this.stack2.pop()
			this.stack1 = this.stack1.push(pop)
		}
	}
	this.stack1 = this.stack1.push(value)
}

func (this *CQueue) DeleteHead1() int {
	var pop = -1
	if !this.stack1.empty() {
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

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
