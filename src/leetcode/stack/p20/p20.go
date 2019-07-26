package p20

type Stack struct {
	Val [1000]int
	Top int
}

func isValid(s string) bool {
	array := []rune(s)
	if len(array)%2 != 0 {
		return false
	}
	for index := 0; index < len(array); index++ {
	}
	return true
}
func match(left, right rune) bool {

}
