package insertionsort

import (
	"learnGo/src/algorithm/sort/tool"
	"testing"
)

func TestSort(t *testing.T) {
	s := new(HalfInsertionSort)
	s.Value = tool.GenerateRandomList()
	tool.SortTool(s)
}
