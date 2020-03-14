package mergesort

import (
	"learnGo/src/algorithm/sort/tool"
	"testing"
)

func TestMergeSortNew_Sort(t *testing.T) {
	s := &MergeSort{}
	s.Value = tool.GenerateRandomList()
	tool.SortTool(s)
}
