package bubble

type BubbleSort [10]int

func (b *BubbleSort) Sort() ([10]int, error) {
	for i := 0; i < len(*b); i++ {
		for j := i + 1; j < len(*b); j++ {
			if b[i] > b[j] {
				b[i], b[j] = b[j], b[i]
			}
		}
	}
	return *b, nil
}
