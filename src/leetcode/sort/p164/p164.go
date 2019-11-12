package p164

func maximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	sort := &QuickSort{
		Value: nums,
	}
	sort.Sort()
	max := 0
	for index := 1; index < len(sort.Value); index++ {
		gap := sort.Value[index] - sort.Value[index-1]
		if gap > max {
			max = gap
		}
	}
	return max
}

type QuickSort struct {
	Value []int
}

func (q *QuickSort) Sort() {
	if len(q.Value) <= 1 {
		return
	}
	quickSort(q.Value, 0, len(q.Value)-1)
}

func quickSort(s []int, left, right int) {
	temp, index := s[left], left
	i, j := left, right
	for j >= i {
		// 比temp大的放在右边
		for j >= index && s[j] >= temp {
			j--
		}
		if j >= index {
			s[index] = s[j]
			index = j
		}

		// 比temp小的放在左边
		for index >= i && s[i] <= temp {
			i++
		}
		if index >= i {
			s[index] = s[i]
			index = i
		}
	}
	s[index] = temp
	if index-left > 1 {
		quickSort(s, left, index-1)
	}
	if right-index > 1 {
		quickSort(s, index+1, right)
	}
}
