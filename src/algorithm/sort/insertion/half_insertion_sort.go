package insertionsort

type HalfInsertionSort struct {
	Value []int
}

func (s *HalfInsertionSort) Sort() ([]int, error) {
	if len(s.Value) <= 1 {
		return s.Value, nil
	}
	for i := 1; i < len(s.Value); i++ {
		// 如果待排序值小于前值，说明需要排序
		if s.Value[i] < s.Value[i-1] {
			key, low, high := s.Value[i], 0, i
			for low <= high {
				mid := (low + high) / 2
				if s.Value[mid] > key {
					high = mid - 1
				} else if s.Value[mid] < key {
					low = mid + 1
				} else {
					high = mid
					break
				}
			}
			var k int
			// 循环结束后，high 的值便是 key 需要在待排序中的位置
			for k = i - 1; k > high; k-- {
				s.Value[k+1] = s.Value[k]
			}
			s.Value[k+1] = key
		}
	}
	return s.Value, nil
}
