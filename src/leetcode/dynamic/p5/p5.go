package p5

func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen == 0 {
		return ""
	}
	left, len := 0, 1
	dp := make([][]int, sLen)
	for i := range dp {
		dp[i] = make([]int, sLen)
	}
	for i := 0; i < sLen; i++ {
		dp[i][i] = 1
		for j := 0; j < i; j++ {
			if (s[i] == s[j]) && (i-j < 2 || dp[j+1][i-1] == 1) {
				dp[j][i] = 1
			}
			if dp[j][i] == 1 && len < i-j+1 {
				len = i - j + 1
				left = j
			}
		}
	}
	return s[left : left+len]
}
