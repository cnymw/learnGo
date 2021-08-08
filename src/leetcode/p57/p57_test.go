package p57

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	intervals = insert(intervals, newInterval)
	fmt.Printf("%v\n", intervals)
}
