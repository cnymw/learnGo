package p4

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		mid := totalLen / 2
		return float64(findKthElement(nums1, nums2, mid+1))
	} else {
		mid1, mid2 := totalLen/2-1, totalLen/2
		return float64(findKthElement(nums1, nums2, mid1+1)+findKthElement(nums1, nums2, mid2+1)) / 2.0
	}
}

func findKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}

		if k == 1 {
			return int(math.Min(float64(nums1[index1]), float64(nums2[index2])))
		}

		half := k / 2
		newIndex1 := int(math.Min(float64(index1+half), float64(len(nums1)))) - 1
		newIndex2 := int(math.Min(float64(index2+half), float64(len(nums2)))) - 1

		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]
		if pivot1 > pivot2 {
			k = k - (newIndex2 - index2) - 1
			index2 = newIndex2 + 1
		} else {
			k = k - (newIndex1 - index1) - 1
			index1 = newIndex1 + 1
		}
	}
}
