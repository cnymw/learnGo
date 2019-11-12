package p115

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	dp := make([][]int, len(t))
	for i := range dp {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			if i > j {
				dp[i][j] = 0
			} else if s[j] == t[i] {
				if i == 0 && j == 0 {
					dp[i][j] = 1
				} else if i == 0 {
					dp[i][j] = dp[i][j-1] + 1
				} else if j == 0 {
					dp[i][j] = dp[i-1][j] + 1
				} else {
					dp[i][j] = dp[i][j-1] + dp[i-1][j-1]
				}
			} else if s[j] != t[i] {
				if j == 0 {
					dp[i][j] = 0
				} else {
					dp[i][j] = dp[i][j-1]
				}

			}
		}
	}
	return dp[len(t)-1][len(s)-1]
}
