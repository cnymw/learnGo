package p30

import (
	"fmt"
	"testing"
)

func TestFindSubstring(t *testing.T) {
	fmt.Println(findSubstring("barfoothefoobarman", []string{"foo", "bar"}))
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "word"}))
}

func TestFindSubstring1(t *testing.T) {
	fmt.Println(findSubstring("wordgoodgoodgoodbestword", []string{"word", "good", "best", "good"}))
}

func TestFindSubstring2(t *testing.T) {
	fmt.Println(findSubstring("foobarfoobar", []string{"foo", "bar"}))
}

func TestPermutation(t *testing.T) {
	perm := make([]string, 0)
	permutation(&perm, []string{"word", "good", "best", "word"}, 0)
	fmt.Println(perm)
}
