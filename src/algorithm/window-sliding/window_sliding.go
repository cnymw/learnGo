package window_sliding

import (
	"math"
)

// 暴力法
func MaxSumByForce(arr []int, n, k int) (maxSum int) {
	maxSum = 0

	for i := 0; i < n-k+1; i++ {
		currentSum := 0
		for j := 0; j < k; j++ {
			currentSum = currentSum + arr[i+j]
		}
		maxSum = int(math.Max(float64(currentSum), float64(maxSum)))
	}
	return
}

// 滑动窗口
func MaxSumByWindowSliding(arr []int, n, k int) (maxSum int) {
	maxSum = 0

	for i := 0; i < k; i++ {
		maxSum = maxSum + arr[i]
	}

	windowSum := maxSum
	for i := k; i < n; i++ {
		windowSum = windowSum + arr[i] - arr[i-k]
		maxSum = int(math.Max(float64(windowSum), float64(maxSum)))
	}
	return
}
