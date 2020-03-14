package bubble

import (
	"learnGo/src/algorithm/sort/tool"
	"testing"
)

func TestBubble(t *testing.T) {
	bubble := new(BubbleSort)
	bubble.Value = tool.GenerateRandomList()
	tool.SortTool(bubble)
}
