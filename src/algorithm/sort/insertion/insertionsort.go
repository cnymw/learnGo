package insertionsort

type InsertionSort struct {
	Value []int
}

func (s *InsertionSort) Sort() ([]int, error) {
	if len(s.Value) <= 1 {
		return s.Value, nil
	}
	for i := 1; i < len(s.Value); i++ {
		t, j := s.Value[i], i
		for ; j >= 1 && s.Value[j-1] > t; j-- {
			s.Value[j] = s.Value[j-1]
		}
		s.Value[j] = t
	}
	return s.Value, nil
}
