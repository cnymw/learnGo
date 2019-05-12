package main

import (
	"learnGo/src/algorithm/sort/bubble"
	"learnGo/src/algorithm/sort/tool"
)

func main() {
	var b bubble.BubbleSort
	b, _ = tool.GenerateRandomList()
	tool.SortTool(&b)
}
