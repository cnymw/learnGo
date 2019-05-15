package heapsort

type HeapSort struct {
	Value []int
}

func (h *HeapSort) Sort() ([]int, error) {
	len := len(h.Value)
	for i := len/2 - 1; i >= 0; i-- {
		heap(&h.Value, i, len)
	}

	for ; len > 0; len-- {
		h.Value[0], h.Value[len-1] = h.Value[len-1], h.Value[0]
		heap(&h.Value, 0, len-1)
	}
	return h.Value, nil
}

func heap(value *[]int, index, len int) {
	t, max := index, (*value)[index]
	// 如果左值大于index
	if left(index) < len && (*value)[left(index)] > max {
		t = left(index)
		max = (*value)[left(index)]
	}
	if right(index) < len && (*value)[right(index)] > max {
		t = right(index)
	}
	if index != t {
		(*value)[index], (*value)[t] = (*value)[t], (*value)[index]
		heap(value, t, len)
	}
}

func left(index int) int {
	return 2*index + 1
}

func right(index int) int {
	return 2*index + 2
}
