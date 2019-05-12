package main

import (
	"learnGo/src/algorithm/sort/bubble"
	"learnGo/src/algorithm/sort/select_sort"
	"learnGo/src/algorithm/sort/tool"
)

func main() {
	var bubble bubble.BubbleSort
	bubble, _ = tool.GenerateRandomList()
	tool.SortTool(&bubble)

	var sel selectsort.SelectSort
	sel, _ = tool.GenerateRandomList()
	tool.SortTool(&sel)
}
