package window_sliding

import "math"

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}

	rk, ans := -1, 0

	for i := 0; i < len(s); i++ {
		if i > 0 {
			delete(m, s[i-1])
		}

		for rk+1 < len(s) && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}

		ans = int(math.Max(float64(ans), float64(rk+1-i)))
	}
	return ans
}
