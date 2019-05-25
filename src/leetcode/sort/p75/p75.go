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
