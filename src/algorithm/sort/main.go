package main

import (
	"learnGo/src/algorithm/sort/bubble"
	"learnGo/src/algorithm/sort/insertion"
	"learnGo/src/algorithm/sort/mergesort"
	"learnGo/src/algorithm/sort/selectsort"
	"learnGo/src/algorithm/sort/shell"
	"learnGo/src/algorithm/sort/tool"
)

func main() {
	var bubble bubble.BubbleSort
	bubble, _ = tool.GenerateRandomList()
	tool.SortTool(&bubble)

	var sel selectsort.SelectSort
	sel, _ = tool.GenerateRandomList()
	tool.SortTool(&sel)

	var insert insertionsort.InsertionSort
	insert, _ = tool.GenerateRandomList()
	tool.SortTool(&insert)

	var shell shellsort.ShellSort
	shell, _ = tool.GenerateRandomList()
	tool.SortTool(&shell)

	var merge mergesort.MergeSort
	merge, _ = tool.GenerateRandomList()
	tool.SortTool(&merge)
}
