package hashmap

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return nil
	}
	// wLen 是 words 的个数，wordLen 是单个 word 的长度
	wLen, wordLen := len(words), len(words[0])

	// words总长为 wLen*wordLen
	result, wordsLen := make([]int, 0), wLen*wordLen
	m := make(map[string]int)
	for _, w := range words {
		m[w]++
	}

	for index := 0; index+wordsLen <= len(s); index++ {
		ich := s[index : index+wordLen]

		if _, ok := m[ich]; ok {
			rCnt := make(map[string]int)
			var j int
			for j = 0; j < wLen; j++ {
				jch := s[index+j*wordLen : index+(j+1)*wordLen]
				rCnt[jch]++
				if rCnt[jch] > m[jch] {
					break
				}
			}
			if j == wLen {
				result = append(result, index)
			}
		}
	}

	return result
}
