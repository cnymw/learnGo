package bubble

type BubbleSort struct {
	Value []int
}

func (b *BubbleSort) Sort() ([]int, error) {
	for i := 0; i < len(b.Value); i++ {
		for j := i + 1; j < len(b.Value); j++ {
			if b.Value[i] > b.Value[j] {
				b.Value[i], b.Value[j] = b.Value[j], b.Value[i]
			}
		}
	}
	return b.Value, nil
}
