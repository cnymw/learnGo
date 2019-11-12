package p3

func lengthOfLongestSubstring(s string) int {
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
