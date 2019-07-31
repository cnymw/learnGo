package p150

import (
	"fmt"
	"strings"
	"testing"
)

func TestMiddleToBack(t *testing.T) {
	tokens := "1+2*3+2"
	output := middleToBack(strings.Split(tokens, ""))
	fmt.Println(output)
}

func TestMiddleToBack2(t *testing.T) {
	tokens := "(1+2)*(2+3)"
	output := middleToBack(strings.Split(tokens, ""))
	fmt.Println(output)
}

func TestMiddleToBack3(t *testing.T) {
	tokens := "1+2*3-(4+5)"
	output := middleToBack(strings.Split(tokens, ""))
	fmt.Println(output)
}
