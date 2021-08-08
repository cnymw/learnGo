package p57

import (
	"math"
	"sort"
)

type Interval struct {
	Value [][]int
}

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	interval := new(Interval)
	interval.Value = intervals
	sort.Sort(interval)

	result := make([][]int, 0)
	result = append(result, interval.Value[0])
	for i := 1; i < len(interval.Value); i++ {
		if result[end(result)][1] < interval.Value[i][0] {
			result = append(result, interval.Value[i])
		} else {
			result[end(result)][1] = int(math.Max(float64(result[end(result)][1]), float64(interval.Value[i][1])))
		}
	}
	return result
}

func end(i [][]int) int {
	return len(i) - 1
}

func (interval *Interval) Len() int {
	return len(interval.Value)
}

func (interval *Interval) Less(i, j int) bool {
	return interval.Value[i][0] < interval.Value[j][0]
}

func (interval *Interval) Swap(i, j int) {
	interval.Value[i], interval.Value[j] = interval.Value[j], interval.Value[i]
}
