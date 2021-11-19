package p11

import "math"

func maxArea(height []int) int {
	l, r, area := 0, len(height)-1, 0
	for l < r {
		a := int(math.Min(float64(height[l]), float64(height[r])) * float64(r-l))
		area = int(math.Max(float64(area), float64(a)))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return area
}
