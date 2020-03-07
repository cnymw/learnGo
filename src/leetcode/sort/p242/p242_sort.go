package p242

import (
	"bytes"
	"sort"
)

type Bytes []byte

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sBytes, tBytes := Bytes(s), Bytes(t)
	sort.Sort(sBytes)
	sort.Sort(tBytes)
	return bytes.Equal(sBytes, tBytes)
}

func (s Bytes) Len() int {
	return len(s)
}

func (s Bytes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s Bytes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
