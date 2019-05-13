package selectsort

type SelectSort [10]int

func (s *SelectSort) Sort() ([10]int, error) {
	for i := 0; i < 10; i++ {
		min := i
		for j := i + 1; j < 10; j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
	return *s, nil
}
