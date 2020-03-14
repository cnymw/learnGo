package quicksort

import (
	"learnGo/src/algorithm/sort/tool"
	"testing"
)

func TestQuickSort_Sort(t *testing.T) {
	q := &QuickSortNew{}
	q.Value = tool.GenerateRandomList()
	tool.SortTool(q)
}
