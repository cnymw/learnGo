package heapsort

// 堆排序-算法导论版本
type HeapSortAlgorithms struct {
	Value []int
}

func (s *HeapSortAlgorithms) Sort() ([]int, error) {
	s.buildMaxHeap()
	for i := len(s.Value) - 1; i > 0; i-- {
		s.Value[i], s.Value[0] = s.Value[0], s.Value[i]
		s.maxHeapify(0, i)
	}
	return s.Value, nil
}

func (s *HeapSortAlgorithms) maxHeapify(t, len int) {
	if t >= len {
		return
	}
	largest, left, right := t, Left(t), Right(t)

	if left < len && s.Value[left] > s.Value[largest] {
		largest = left
	}

	if right < len && s.Value[right] > s.Value[largest] {
		largest = right
	}

	if largest != t {
		s.Value[t], s.Value[largest] = s.Value[largest], s.Value[t]
		s.maxHeapify(largest, len)
	}
}

func (s *HeapSortAlgorithms) buildMaxHeap() {
	for i := len(s.Value)/2 - 1; i >= 0; i-- {
		s.maxHeapify(i, len(s.Value))
	}
}

func Left(t int) int {
	return 2*t + 1
}

func Right(t int) int {
	return 2*t + 2
}
