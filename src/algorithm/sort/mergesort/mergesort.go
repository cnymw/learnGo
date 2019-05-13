package mergesort

type MergeSort [10]int

func (m *MergeSort) Sort() ([10]int, error) {
	s := make([]int, 0)
	for _, v := range *m {
		s = append(s, v)
	}
	s = mergeSort(s)
	for i, v := range s {
		m[i] = v
	}
	return *m, nil
}

func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}

func merge(left, right []int) (result []int) {
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
