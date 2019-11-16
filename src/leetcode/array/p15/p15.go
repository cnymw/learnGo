package p15

import "sort"

type Interval struct {
	Value []int
}

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if nums == nil || len(nums) < 3 {
		return result
	}
	interval := new(Interval)
	interval.Value = nums
	sort.Sort(interval)

	for index := 0; index < interval.Len(); index++ {
		if interval.Value[index] > 0 {
			break
		}
		if index >= 1 && interval.Value[index] == interval.Value[index-1] {
			continue
		}
		for L, R := index+1, interval.Len()-1; L < R; {
			sum := interval.Value[index] + interval.Value[L] + interval.Value[R]
			if sum == 0 {
				tmp := []int{interval.Value[index], interval.Value[L], interval.Value[R]}
				result = append(result, tmp)
				for ; L+1 < interval.Len() && interval.Value[L] == interval.Value[L+1]; L++ {

				}
				for ; R-1 >= 0 && interval.Value[R] == interval.Value[R-1]; R-- {

				}
				L++
				R--
				continue
			} else if sum > 0 {
				R--
				continue
			} else {
				L++
				continue
			}
		}
	}
	return result
}

func (interval *Interval) Len() int {
	return len(interval.Value)
}

func (interval *Interval) Less(i, j int) bool {
	return interval.Value[i] < interval.Value[j]
}

func (interval *Interval) Swap(i, j int) {
	interval.Value[i], interval.Value[j] = interval.Value[j], interval.Value[i]
}
