package quicksort

type QuickSort struct {
	Value []int
}

func (q *QuickSort) Sort() ([]int, error) {
	if len(q.Value) <= 1 {
		return nil, nil
	}
	quickSort(q.Value, 0, len(q.Value)-1)
	return q.Value, nil
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
