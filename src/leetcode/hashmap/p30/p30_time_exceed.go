package p30

// 超出时间限制
func findSubstring1(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	wordLen := len(words[0]) * len(words)
	sMap := make(map[string][]int, 0)
	for index := 0; index+wordLen <= len(s); index++ {
		if list, ok := sMap[s[index:index+wordLen]]; ok {
			list = append(list, index)
			sMap[s[index:index+wordLen]] = list
		} else {
			list := make([]int, 0)
			list = append(list, index)
			sMap[s[index:index+wordLen]] = list
		}
	}
	perm := make([]string, 0)
	permutation(&perm, words, 0)
	rMap := make(map[int]bool, 0)
	for _, p := range perm {
		if list, ok := sMap[p]; ok {
			for _, l := range list {
				rMap[l] = true
			}
		}
	}
	result := make([]int, 0)
	for k := range rMap {
		result = append(result, k)
	}
	return result
}

// 获取 words 全排列数据
func permutation(perm *[]string, words []string, i int) {
	if len(words) == 0 || i < 0 || i > len(words) {
		return
	}
	if i == len(words)-1 {
		var temp string
		for _, w := range words {
			temp += w
		}
		*perm = append(*perm, temp)
	} else {
		for j := i; j < len(words); j++ {
			temp := words[j]
			words[j] = words[i]
			words[i] = temp
			permutation(perm, words, i+1)
			temp = words[j]
			words[j] = words[i]
			words[i] = temp
		}
	}
}
