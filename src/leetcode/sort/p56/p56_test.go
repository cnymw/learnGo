package p56

import (
	"fmt"
	"testing"
)

func TestMerge1(t *testing.T) {
	var intervals [][]int
	intervals = [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	intervals = merge(intervals)
	fmt.Printf("%v\n", intervals)

	intervals = [][]int{{1, 4}, {4, 5}}
	intervals = merge(intervals)
	fmt.Printf("%v\n", intervals)

	intervals = [][]int{{1, 3}}
	intervals = merge(intervals)
	fmt.Printf("%v\n", intervals)

	intervals = [][]int{{1, 4}, {1, 5}}
	intervals = merge(intervals)
	fmt.Printf("%v\n", intervals)

	intervals = [][]int{{1, 4}, {2, 3}}
	intervals = merge(intervals)
	fmt.Printf("%v\n", intervals)

	intervals = [][]int{{1, 4}, {0, 2}, {3, 5}}
	intervals = merge(intervals)
	fmt.Printf("%v\n", intervals)

	intervals = [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	intervals = merge(intervals)
	fmt.Printf("%v\n", intervals)
}
