package better

func isValid(s string) bool {
	array := []rune(s)
	if len(array)%2 != 0 {
		return false
	}

	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	stack := make(Stack, 0)
	var pop rune
	for _, r := range array {
		_, ok := mapping[r]
		if ok {
			stack, pop = stack.Pop()
			if mapping[r] != pop {
				return false
			}
		} else {
			stack = stack.Push(r)
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

type Stack []rune

func (s Stack) Push(v rune) Stack {
	return append(s, v)
}

func (s Stack) Pop() (Stack, rune) {
	if len(s) == 0 {
		return nil, 0
	}
	return s[:len(s)-1], s[len(s)-1]
}
