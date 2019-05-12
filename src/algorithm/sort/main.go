package main

import (
	"learnGo/src/algorithm/sort/bubble"
	"learnGo/src/algorithm/sort/insertion_sort"
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

	var insert insertionsort.InsertionSort
	//insert,_ = tool.GenerateRandomList()
	insert = [10]int{94, 44, 42, 45, 48, 20, 43, 51, 89, 8}
	tool.SortTool(&insert)
}
