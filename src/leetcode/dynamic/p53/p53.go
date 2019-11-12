package p53

import (
	"math"
)

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]

	for index := 1; index < len(nums); index++ {
		dp[index] = int(math.Max(float64(dp[index-1]+nums[index]), float64(nums[index])))
	}

	max := dp[0]
	for index := 1; index < len(nums); index++ {
		if dp[index] > max {
			max = dp[index]
		}
	}
	return max
}
