package p27

func removeElement(nums []int, val int) int {
	for index := 0; index < len(nums); {
		if nums[index] == val {
			nums = append(nums[:index], nums[index+1:]...)
		} else {
			index++
		}
	}
	return len(nums)
}
