package window_sliding

import (
	"fmt"
	"testing"
)

func TestMaxSumByForce(t *testing.T) {
	arr := []int{1, 4, 2, 10, 2, 3, 1, 0, 20}

	fmt.Println(MaxSumByForce(arr, len(arr), 4))
}

func TestMaxSumByWindowSliding(t *testing.T) {
	arr := []int{1, 4, 2, 10, 2, 3, 1, 0, 20}

	fmt.Println(MaxSumByWindowSliding(arr, len(arr), 4))
}
