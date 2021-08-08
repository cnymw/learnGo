package p225

type Element interface{}

type Queue interface {
	Offer(e Element)
	Poll() Element
	Clear() bool
	Size() int
	IsEmpty() bool
}

type sliceEntry struct {
	element []Element
}

func NewQueue() *sliceEntry {
	return &sliceEntry{}
}

func (entry *sliceEntry) Offer(e Element) {
	entry.element = append(entry.element, e)
}

func (entry *sliceEntry) Poll() Element {
	if entry.IsEmpty() {
		return nil
	}
	first := entry.element[0]
	entry.element = entry.element[1:]
	return first
}

func (entry *sliceEntry) Clear() bool {
	if entry.IsEmpty() {
		return false
	}
	for i := 0; i < len(entry.element); i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}
func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry) IsEmpty() bool {
	return entry.Size() == 0
}

type MyStack struct {
	queueA *sliceEntry
	queueB *sliceEntry
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	qa, qb := NewQueue(), NewQueue()
	return MyStack{queueA: qa, queueB: qb}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	if this.queueA.Size() != 0 {
		this.queueA.Offer(x)
	} else if this.queueB.Size() != 0 {
		this.queueB.Offer(x)
	} else {
		this.queueA.Offer(x)
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.queueA.Size() == 0 && this.queueB.Size() == 0 {
		return 0
	}
	result := 0
	if this.queueA.Size() != 0 {
		for this.queueA.Size() > 0 {
			tmp := this.queueA.Poll()
			if this.queueA.Size() != 0 {
				this.queueB.Offer(tmp)
			} else {
				result = tmp.(int)
			}
		}
	} else {
		for this.queueB.Size() > 0 {
			tmp := this.queueB.Poll()
			if this.queueB.Size() != 0 {
				this.queueA.Offer(tmp)
			} else {
				result = tmp.(int)
			}
		}
	}
	return result
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.queueA.Size() != 0 {
		for this.queueA.Size() > 0 {
			tmp := this.queueA.Poll()
			if this.queueA.Size() != 0 {
				this.queueB.Offer(tmp)
			} else {
				this.queueB.Offer(tmp)
				return tmp.(int)
			}
		}
	} else {
		for this.queueB.Size() > 0 {
			tmp := this.queueB.Poll()
			if this.queueB.Size() != 0 {
				this.queueA.Offer(tmp)
			} else {
				this.queueA.Offer(tmp)
				return tmp.(int)
			}
		}
	}
	return 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.queueA.Size() != 0 || this.queueB.Size() != 0 {
		return false
	}
	return true
}
