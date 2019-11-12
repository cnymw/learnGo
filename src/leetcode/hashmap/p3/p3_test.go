package p3

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))  //3
	fmt.Println(lengthOfLongestSubstring("bbbbb"))     //1
	fmt.Println(lengthOfLongestSubstring("pwwkew"))    //3
	fmt.Println(lengthOfLongestSubstring("dvdf"))      //3
	fmt.Println(lengthOfLongestSubstring("tmmzuxt"))   //5
	fmt.Println(lengthOfLongestSubstring("bpfbhmipx")) //7
}
