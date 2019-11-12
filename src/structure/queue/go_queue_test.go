package queue

import (
	"fmt"
	"testing"
)

func TestQueue_Dequeue(t *testing.T) {
	queue := NewQueue(10)
	queue.Enqueue(111)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
	queue.Enqueue(6)
	queue.Enqueue(7)
	queue.Enqueue(8)
	queue.Enqueue(9)
	queue.Enqueue(10)

	for _, v := range queue.Q {
		fmt.Println(v)
	}
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Head)
	fmt.Println(queue.Tail)
}
