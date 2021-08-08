package p1030

import (
	"math"
)

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	m, max := make(map[int][][]int, 0), 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			tmp := []int{i, j}
			distance := int(math.Abs(float64(r0-i)) + math.Abs(float64(c0-j)))
			result, ok := m[distance]
			if ok {
				result = append(result, tmp)
				m[distance] = result
			} else {
				result := make([][]int, 0)
				result = append(result, tmp)
				m[distance] = result
			}
			if distance > max {
				max = distance
			}
		}
	}

	result := make([][]int, 0)
	for i := 0; i <= max; i++ {
		tmp := m[i]
		result = append(result, tmp...)
	}
	return result
}
