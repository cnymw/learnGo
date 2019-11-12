package p242

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sCh, tCh := make([]int, 0), make([]int, 0)
	for _, ch := range s {
		sCh = append(sCh, int(ch))
	}
	for _, ch := range t {
		tCh = append(tCh, int(ch))
	}
	sSort := QuickSort{
		Value: sCh,
	}
	tSort := QuickSort{
		Value: tCh,
	}
	sSort.Sort()
	tSort.Sort()
	for index := 0; index < len(sSort.Value); index++ {
		if sSort.Value[index] != tSort.Value[index] {
			return false
		}
	}
	return true
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
