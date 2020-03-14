package quicksort

type QuickSortNew struct {
	Value []int
}

func (q *QuickSortNew) Sort() ([]int, error) {
	quickSortNew(q.Value, 0, len(q.Value)-1)
	return q.Value, nil
}

func quickSortNew(s []int, l, r int) {
	if l < r {
		q := patitionNew(s, l, r)
		quickSortNew(s, l, q-1)
		quickSortNew(s, q+1, r)
	}

}
func patitionNew(s []int, l, r int) int {
	key, index := s[r], l-1

	for j := l; j < r; j++ {
		if s[j] <= key {
			index++
			s[index], s[j] = s[j], s[index]
		}
	}
	s[index+1], s[r] = s[r], s[index+1]
	return index + 1
}
