package p155

type MinStack struct {
	top    *node
	length int
}
type node struct {
	value int
	prev  *node
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{nil, 0}
}

func (this *MinStack) Push(x int) {
	n := &node{x, this.top}
	this.top = n
	this.length++
}

func (this *MinStack) Pop() {
	if this.length == 0 {
		return
	}

	n := this.top
	this.top = n.prev
	this.length--
}

func (this *MinStack) Top() int {
	if this.length == 0 {
		return 0
	}
	return this.top.value
}

func (this *MinStack) GetMin() int {
	cur := this.top
	min := cur.value
	for cur != nil {
		if min > cur.value {
			min = cur.value
		}
		cur = cur.prev
	}
	return min
}
