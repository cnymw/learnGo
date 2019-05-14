package p56

import (
	"fmt"
	"testing"
)

func TestMerge1(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	//intervals = merge(intervals)
	//fmt.Printf("%v\n",intervals)
	//
	//intervals = [][]int{{1,4},{4,5}}
	//intervals = merge(intervals)
	//fmt.Printf("%v\n",intervals)
	//
	//intervals = [][]int{{1,3}}
	//intervals = merge(intervals)
	//fmt.Printf("%v\n",intervals)
	//
	//intervals = [][]int{{1,4},{1,5}}
	//intervals = merge(intervals)
	//fmt.Printf("%v\n",intervals)
	//
	//intervals = [][]int{{1,4},{2,3}}
	//intervals = merge(intervals)
	//fmt.Printf("%v\n",intervals)

	intervals = [][]int{{1, 4}, {0, 2}, {3, 5}}
	intervals = merge(intervals)
	fmt.Printf("%v\n", intervals)
}
