package radixsort

type RadixSort struct {
	Value []int
}

func (r *RadixSort) Sort() ([]int, error) {
	temp := make([][]int, 10)
	for i := 0; i < 10; i++ {
		temp[i] = make([]int, len(r.Value))
	}
	count := make([]int, 10)
	maxLoop, loop, n := 3, 1, 1
	for ; loop <= maxLoop; loop++ {
		k := 0
		for _, v := range r.Value {
			lsd := (v / n) % 10
			temp[lsd][count[lsd]] = v
			count[lsd]++
		}
		for i := 0; i < 10; i++ {
			if count[i] != 0 {
				for j := 0; j < count[i]; j++ {
					r.Value[k] = temp[i][j]
					k++
				}
			}
			count[i] = 0
		}
		n *= 10
	}
	return r.Value, nil
}
