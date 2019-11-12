package p70

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	oneStep, twoStep, total := 1, 2, 0
	for index := 3; index <= n; index++ {
		total = oneStep + twoStep
		oneStep = twoStep
		twoStep = total
	}
	return total
}
