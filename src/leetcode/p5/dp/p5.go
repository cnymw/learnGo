package dp

func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return ""
	}
	left, maxlen := 0, 1

	dp := make([][]int, sLen)
	for i := 0; i < sLen; i++ {
		dp[i] = make([]int, sLen)
	}

	for j := 0; j < sLen; j++ {
		dp[j][j] = 1
		for i := 0; i < j; i++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1] == 1) {
				dp[i][j] = 1
			}
			if dp[i][j] == 1 && maxlen < j-i+1 {
				left = i
				maxlen = j - i + 1
			}
		}
	}

	return s[left : left+maxlen]
}
