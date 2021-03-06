package heapsort

import (
	"learnGo/src/algorithm/sort/tool"
	"testing"
)

func TestHeapSortAlgorithms_Sort(t *testing.T) {
	s := new(HeapSortAlgorithms)
	s.Value = tool.GenerateRandomList()
	tool.SortTool(s)
}
