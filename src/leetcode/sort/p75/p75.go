package p75

func sortColors(nums []int) {
	count := make([]int, 3)
	for _, v := range nums {
		count[v]++
	}
	index := 0
	for i := 0; i < 3; i++ {
		if count[i] > 0 {
			for j := 0; j < count[i]; j++ {
				nums[index] = i
				index++
			}
		}
	}
}

// 三指针法
func sortColorsWithPointer(nums []int) {
	index, l, r := 0, 0, len(nums)-1
	for index <= r {
		if index < l || nums[index] == 1 {
			index++
		} else if nums[index] == 2 {
			nums[index], nums[r] = nums[r], nums[index]
			r--
		} else if nums[index] == 0 {
			nums[index], nums[l] = nums[l], nums[index]
			l++
		}
	}
}
