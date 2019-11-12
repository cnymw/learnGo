package p1

func twoSum(nums []int, target int) []int {
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
