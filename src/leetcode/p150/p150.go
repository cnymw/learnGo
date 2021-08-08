package p150

import (
	"regexp"
	"strconv"
)

var pattern = "\\d+"

func evalInterface(v interface{}) int {
	if i, ok := v.(string); ok {
		r, _ := strconv.Atoi(i)
		return r
	} else {
		return v.(int)
	}
}
func evalRPN(tokens []string) int {
	s := New()
	for _, t := range tokens {
		if isDigit(t) {
			s.Push(t)
		} else {
			one, two := evalInterface(s.Pop()), evalInterface(s.Pop())
			result := 0
			if t == "+" {
				result = two + one
			} else if t == "-" {
				result = two - one
			} else if t == "*" {
				result = two * one
			} else if t == "/" {
				result = two / one
			}
			s.Push(result)
		}
	}
	result := evalInterface(s.Pop())
	return result
}

func isDigit(s string) bool {
	m, _ := regexp.MatchString(pattern, s)
	return m
}

type (
	Stack struct {
		top    *node
		length int
	}
	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new stack
func New() *Stack {
	return &Stack{nil, 0}
}

// Return the number of items in the stack
func (this *Stack) Len() int {
	return this.length
}

// View the top item on the stack
func (this *Stack) Peek() interface{} {
	if this.length == 0 {
		return nil
	}
	return this.top.value
}

// Pop the top item of the stack and return it
func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

// Push a value onto the top of the stack
func (this *Stack) Push(value interface{}) {
	n := &node{value, this.top}
	this.top = n
	this.length++
}

func (this *Stack) Empty() bool {
	return this.length == 0
}
