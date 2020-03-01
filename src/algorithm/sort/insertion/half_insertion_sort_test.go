package insertionsort

import (
	"learnGo/src/algorithm/sort/tool"
	"testing"
)

func TestSort(t *testing.T) {
	s := new(HalfInsertionSort)
	s.Value, _ = tool.GenerateRandomList()
	tool.SortTool(s)
}
