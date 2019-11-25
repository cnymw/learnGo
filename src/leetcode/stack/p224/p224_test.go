package p224

import (
	"fmt"
	"testing"
)

func TestCalculate(t *testing.T) {
	//fmt.Println(calculate("1 + 1"))
	//fmt.Println(calculate("2-1 + 2 "))
	fmt.Println(calculate("(1+(4+5+2)-3)+(6+8)"))
}

func TestCalculate1(t *testing.T) {
	fmt.Println(calculate("2147483647"))
}

func TestCalculate2(t *testing.T) {
	fmt.Println(calculate("(1-(3-4))"))
}

func TestCalculate3(t *testing.T) {
	fmt.Println(calculate(" 2-1 + 2 "))
}
