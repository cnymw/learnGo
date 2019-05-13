package selectsort

type SelectSort struct {
	Value []int
}

func (s *SelectSort) Sort() ([]int, error) {
	for i := 0; i < 10; i++ {
		min := i
		for j := i + 1; j < 10; j++ {
			if s.Value[j] < s.Value[min] {
				min = j
			}
		}
		s.Value[i], s.Value[min] = s.Value[min], s.Value[i]
	}
	return s.Value, nil
}
