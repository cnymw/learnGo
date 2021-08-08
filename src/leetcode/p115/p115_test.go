package p115

import (
	"fmt"
	"testing"
)

func TestNumDistinct1(t *testing.T) {
	fmt.Println(numDistinct("rabbbit", "rabbit"))
}

func TestNumDistinct2(t *testing.T) {
	fmt.Println(numDistinct("babgbag", "bag"))
}

func TestNumDistinct3(t *testing.T) {
	fmt.Println(numDistinct("aabb", "abb"))
}

func TestNumDistinct4(t *testing.T) {
	fmt.Println(numDistinct("b", "a"))
}
