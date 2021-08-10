package p4

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1, nums2 := []int{1, 2}, []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
