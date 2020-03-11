package p976

import (
	"sort"
)

type Arr []int

func largestPerimeter(A []int) int {
	s := Arr(A)

	sort.Sort(s)

	for i := 0; i <= len(s)-3; i++ {
		if judgeTriangle(s[i], s[i+1], s[i+2]) {
			return s[i] + s[i+1] + s[i+2]
		}
	}
	return 0
}

func judgeTriangle(a, b, c int) bool {
	if a < b+c && a-b < c && a-c < b {
		return true
	} else {
		return false
	}
}

func (s Arr) Len() int {
	return len(s)
}

func (s Arr) Less(i, j int) bool {
	return s[i] > s[j]
}

func (s Arr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
