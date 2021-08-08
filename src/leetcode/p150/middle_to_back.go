package p150

import "strings"

// 中缀表达式转后缀表达式
// 中：1+2*3-(4+5)
// ->
// 后：123*+45+-
func middleToBack(tokens []string) string {
	output, st := make([]string, 0), New()
	for _, v := range tokens {
		if isDigit(v) {
			output = append(output, v)
		} else if v == "(" {
			st.Push(v)
		} else if v == ")" {
			for {
				tmp := st.Pop().(string)
				if tmp == "(" {
					break
				}
				output = append(output, tmp)
			}
		} else {
			if v == "-" || v == "+" {

				for {
					if !st.Empty() {
						tmp := st.Peek().(string)
						if tmp == "(" {
							break
						}
						tmp = st.Pop().(string)
						output = append(output, tmp)
					} else {
						break
					}
				}

				st.Push(v)
			} else if v == "*" || v == "/" {

				for {
					if !st.Empty() {
						tmp := st.Peek().(string)
						if tmp == "(" || tmp == "+" || tmp == "-" {
							break
						}
						tmp = st.Pop().(string)
						output = append(output, tmp)
					} else {
						break
					}
				}
				st.Push(v)
			}
		}
	}
	for !st.Empty() {
		output = append(output, st.Pop().(string))
	}
	return strings.Join(output, "")
}
