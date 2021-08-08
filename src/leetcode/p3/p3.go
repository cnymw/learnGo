package p3

import "math"

// 哈希表
func lengthOfLongestSubstringByHashMap(s string) int {
	if len(s) == 0 {
		return 0
	}
	m, length := make(map[rune]int, 0), 1
	sub := make([]int32, 0)
	runes := []rune(s)
	for index, r := range s {
		if i, ok := m[r]; ok {
			m = make(map[rune]int, 0)
			sub = runes[i+1 : index]
			sub = append(sub, r)

			for j, k := index-len(sub)+1, 0; j <= index && k < len(sub); {
				m[sub[k]] = j
				k++
				j++
			}
		} else {
			sub = append(sub, r)
			if len(sub) > length {
				length = len(sub)
			}
			m[r] = index
		}
	}
	return length
}

// 滑动窗口
func lengthOfLongestSubstringByWindowSliding(s string) int {
	n := len(s)
	m := map[byte]int{}

	rk, ans := -1, 0
	for i := 0; i < n; i++ {
		// 滑动到下一个节点，将前一个节点删除
		if i > 0 {
			delete(m, s[i-1])
		}

		// 如果满足条件(无重复字符)，那么会持续往后面滑动
		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}

		ans = int(math.Max(float64(ans), float64(rk+1-i)))
	}

	return ans
}
