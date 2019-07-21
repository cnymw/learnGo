package p75

import (
	"fmt"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 1, 2, 0, 1, 0}
	sortColors(nums)
	for _, v := range nums {
		fmt.Println(v)
	}

	nums2 := []int{2, 1, 2, 0, 1, 0}
	sortColorsWithPointer(nums2)
	for _, v := range nums2 {
		fmt.Println(v)
	}
	// 1. index = 0 , l = 0 , r = 5
	// 1. nums[index]==2 -> nums[0]=2,nums[5]=2 -> 0,1,2,0,1,2 -> r--=4
	// 2. index = 0 , l = 0 , r = 4
	// 2. nums[index]==0 -> nums[0]=0,nums[0]=0 -> 0,1,2,0,1,2 -> l++=1
	// 3. index = 0 , l = 1 , r = 4
	// 3. index<l -> index++=1
	// 4. index = 1 , l = 1 , r = 4
	// 4. nums[index]==1 -> index++=2

}
