package p3

import (
	"fmt"
	"testing"
)

func TestLengthOfLongestSubstringByHashMap(t *testing.T) {
	fmt.Println(lengthOfLongestSubstringByHashMap("abcabcbb"))  //3
	fmt.Println(lengthOfLongestSubstringByHashMap("bbbbb"))     //1
	fmt.Println(lengthOfLongestSubstringByHashMap("pwwkew"))    //3
	fmt.Println(lengthOfLongestSubstringByHashMap("dvdf"))      //3
	fmt.Println(lengthOfLongestSubstringByHashMap("tmmzuxt"))   //5
	fmt.Println(lengthOfLongestSubstringByHashMap("bpfbhmipx")) //7
}

func TestLengthOfLongestSubstringByWindowSliding(t *testing.T) {
	fmt.Println(lengthOfLongestSubstringByWindowSliding("abcabcbb"))  //3
	fmt.Println(lengthOfLongestSubstringByWindowSliding("bbbbb"))     //1
	fmt.Println(lengthOfLongestSubstringByWindowSliding("pwwkew"))    //3
	fmt.Println(lengthOfLongestSubstringByWindowSliding("dvdf"))      //3
	fmt.Println(lengthOfLongestSubstringByWindowSliding("tmmzuxt"))   //5
	fmt.Println(lengthOfLongestSubstringByWindowSliding("bpfbhmipx")) //7
}
