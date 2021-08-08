package p349

func intersection(nums1 []int, nums2 []int) []int {
	nums1Map, nums2Map := make(map[int]int), make(map[int]int)
	for _, v := range nums1 {
		nums1Map[v]++
	}
	for _, v := range nums2 {
		if nums1Map[v] > 0 {
			nums2Map[v]++
		}
	}
	result := make([]int, 0)
	for k := range nums2Map {
		result = append(result, k)
	}
	return result
}
