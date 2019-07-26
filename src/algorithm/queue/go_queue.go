package queue

type Queue struct {
	Q    []int
	Head int
	Tail int
}

func NewQueue(size int) *Queue {
	return &Queue{
		Q:    make([]int, size),
		Head: 0,
		Tail: 1,
	}
}

func (queue *Queue) Enqueue(v int) {
	queue.Q[queue.Tail-1] = v
	if queue.Tail == len(queue.Q) {
		queue.Tail = 0
	} else {
		queue.Tail++
	}
}

func (queue *Queue) Dequeue() int {
	v := queue.Q[queue.Head]
	if queue.Head == len(queue.Q) {
		queue.Head = 0
	} else {
		queue.Head++
	}
	return v
}
