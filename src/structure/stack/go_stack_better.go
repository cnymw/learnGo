package stack

type stack []interface{}

func New() stack {
	s := make(stack, 0)
	return s
}
func (s stack) Push(v []interface{}) stack {
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
