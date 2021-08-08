package p242

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
}

func TestIsAnagramHash(t *testing.T) {
	fmt.Println(isAnagramHash("anagram", "nagaram"))
	fmt.Println(isAnagramHash("rat", "car"))
}
