package p42

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}

	fmt.Println(trap(height))
}

func TestNew1(t *testing.T) {
	height := []int{4, 2, 3}

	fmt.Println(trap(height))
}
