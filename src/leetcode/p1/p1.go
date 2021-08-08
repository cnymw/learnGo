package p1

// 方法1：暴力枚举
func twoSumByLoop(nums []int, target int) []int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0}
}

// 方法2：哈希表
func twoSumByHashMap(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))

	for index, n := range nums {
		numsMap[n] = index
	}

	for index, n := range nums {
		m, ok := numsMap[target-n]
		if ok && m != index {
			return []int{index, m}
		}
	}
	return []int{0, 0}
}
