package shellsort

type ShellSort struct {
	Value []int
}

func (s *ShellSort) Sort() ([]int, error) {
	for i := 10 / 3; i >= 1; i-- {
		for j := 0; j < i; j++ {

			// 插入排序
			for m := j + i; m < len(s.Value); m += i {
				t, index := s.Value[m], m
				for ; index >= i && s.Value[index-i] > t; index -= i {
					s.Value[index] = s.Value[index-i]
				}
				s.Value[index] = t
			}
		}
	}
	return s.Value, nil
}
