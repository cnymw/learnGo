package shellsort

type ShellSort [10]int

func (s *ShellSort) Sort() ([10]int, error) {
	for i := 10 / 3; i >= 1; i-- {
		for j := 0; j < i; j++ {

			// 插入排序
			for m := j + i; m < 10; m += i {
				t, index := s[m], m
				for ; index >= i && s[index-i] > t; index -= i {
					s[index] = s[index-i]
				}
				s[index] = t
			}
		}
	}
	return *s, nil
}
