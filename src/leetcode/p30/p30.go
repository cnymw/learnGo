package p30

func findSubstring(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}
	wordsCnt := make(map[string]int, 0)
	for _, word := range words {
		wordsCnt[word]++
	}
	wordsLength := len(words)
	length := len(words[0])
	result := make([]int, 0)
	for index := 0; index+wordsLength*length <= len(s); index++ {
		ch := s[index : index+length]

		if _, ok := wordsCnt[ch]; ok {
			sCnt := make(map[string]int, 0)
			var j int
			for j = 0; j < wordsLength; j++ {
				temp := s[index+j*length : index+(j+1)*length]
				sCnt[temp]++
				if sCnt[temp] > wordsCnt[temp] {
					break
				}
			}
			if j == wordsLength {
				result = append(result, index)
			}
		}
	}
	return result
}
