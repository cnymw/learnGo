package main

import (
	"learnGo/src/algorithm/sort/bubble"
	"learnGo/src/algorithm/sort/tool"
	"testing"
)

func TestBubble(t *testing.T) {
	bubble := new(bubble.BubbleSort)
	bubble.Value, _ = tool.GenerateRandomList()
	tool.SortTool(bubble)
}
