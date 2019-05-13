package insertionsort

type InsertionSort [10]int

func (s *InsertionSort) Sort() ([10]int, error) {
	for i := 1; i < 10; i++ {
		t, j := s[i], i
		for ; j >= 1 && s[j-1] > t; j-- {
			s[j] = s[j-1]
		}
		s[j] = t
	}
	return *s, nil
}
