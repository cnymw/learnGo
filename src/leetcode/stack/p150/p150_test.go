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

func TestEvalRPN(t *testing.T) {
	tokens := []string{"2", "1", "+", "3", "*"}
	fmt.Println(evalRPN(tokens))
}

func TestEvalRPN2(t *testing.T) {
	tokens := []string{"4", "13", "5", "/", "+"}
	fmt.Println(evalRPN(tokens))
}

func TestEvalRPN3(t *testing.T) {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}

func TestEvalRPN4(t *testing.T) {
	tokens := []string{"18"}
	fmt.Println(evalRPN(tokens))
}
