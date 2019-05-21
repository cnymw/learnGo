package bucketsort

import (
	"learnGo/src/algorithm/sort/insertion"
	"math"
)

type BucketSort struct {
	Value []int
}

const bucketSize = 5

func (b *BucketSort) Sort() ([]int, error) {
	min, max := b.Value[0], b.Value[0]
	for _, v := range b.Value {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	bucketCount := int(math.Floor(float64(max-min)/bucketSize) + 1)

	bucket := make([][]int, bucketCount)
	for _, v := range b.Value {
		index := int(math.Floor(float64(v-min) / bucketSize))
		bucket[index] = append(bucket[index], v)
	}
	insert := new(insertionsort.InsertionSort)
	index := 0
	for i := 0; i < bucketCount; i++ {
		insert.Value = bucket[i]
		bucket[i], _ = insert.Sort()
		for _, v := range bucket[i] {
			b.Value[index] = v
			index++
		}
	}
	return b.Value, nil
}
