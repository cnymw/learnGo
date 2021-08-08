package p224

import (
	"regexp"
	"strconv"
)

var pattern = "\\d+"

func calculate(s string) int {
	runes := []rune(s)

	var tokens []string
	stack := New()

	nums := make([]rune, 0)
	var v interface{}
	for _, r := range runes {
		switch r {
		case '(':
			stack = stack.Push(string(r))
		case ')':
			if len(nums) != 0 {
				tokens = append(tokens, string(nums))
				nums = make([]rune, 0)
			}

			for {
				stack, v = stack.Pop()
				if v == "(" || stack.Empty() {
					break
				}
				tokens = append(tokens, v.(string))
			}
		case '+', '-':
			if len(nums) != 0 {
				tokens = append(tokens, string(nums))
				nums = make([]rune, 0)
			}
			if stack.Empty() {
				stack = stack.Push(string(r))
			} else {
				pop := stack.Peek()
				if pop == "(" {
					stack = stack.Push(string(r))
					continue
				} else {
					stack, _ = stack.Pop()
					tokens = append(tokens, pop.(string))
					stack = stack.Push(string(r))
				}
			}
		case ' ':

		default:
			nums = append(nums, r)
		}
	}
	if len(nums) != 0 {
		stack = stack.Push(string(nums))
	}
	for !stack.Empty() {
		stack, v = stack.Pop()
		tokens = append(tokens, v.(string))
	}

	var one, two interface{}
	stack = New()
	for _, t := range tokens {
		if isDigit(t) {
			stack = stack.Push(t)
		} else {
			stack, one = stack.Pop()
			stack, two = stack.Pop()
			one, two := evalInterface(one), evalInterface(two)
			result := 0
			if t == "+" {
				result = two + one
			} else if t == "-" {
				result = two - one
			}
			stack = stack.Push(result)
		}
	}
	stack, v = stack.Pop()
	result := evalInterface(v)
	return result
}

func isDigit(s string) bool {
	m, _ := regexp.MatchString(pattern, s)
	return m
}

func evalInterface(v interface{}) int {
	if i, ok := v.(string); ok {
		r, _ := strconv.Atoi(i)
		return r
	} else {
		return v.(int)
	}
}

type stack []interface{}

func New() stack {
	s := make(stack, 0)
	return s
}
func (s stack) Push(v interface{}) stack {
	return append(s, v)
}

func (s stack) Pop() (stack, interface{}) {
	if s.Len() == 0 {
		return s, nil
	}
	return s[:s.Len()-1], s[s.Len()-1]
}

func (s stack) Len() int {
	return len(s)
}

func (s stack) Peek() interface{} {
	return s[s.Len()-1]
}

func (s stack) Empty() bool {
	return s.Len() == 0
}
