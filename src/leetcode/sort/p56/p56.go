package p56

import "math"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	result := make([][]int, 0)
	temp := make([]int, 0)
	for _, val := range intervals {
		if len(temp) == 0 {
			temp = val
		} else {
			if (temp[0] >= val[0] && temp[0] <= val[1]) || (val[0] >= temp[0] && val[0] <= temp[1]) {
				t := []int{int(math.Min(float64(temp[0]), float64(val[0]))), int(math.Max(float64(temp[1]), float64(val[1])))}
				result = append(result, t)
				temp = temp[0:0]
			} else {
				result = append(result, temp)
				result = append(result, val)
			}
		}
	}
	return result
}
