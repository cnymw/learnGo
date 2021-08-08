package p242

func isAnagramHash(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counter := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]-'a']++
		counter[t[i]-'a']--
	}
	for i := range counter {
		if counter[i] != 0 {
			return false
		}
	}
	return true
}
