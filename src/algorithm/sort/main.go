package main

import (
	"learnGo/src/algorithm/sort/bubble"
	"learnGo/src/algorithm/sort/heapsort"
	"learnGo/src/algorithm/sort/insertion"
	"learnGo/src/algorithm/sort/mergesort"
	"learnGo/src/algorithm/sort/quicksort"
	"learnGo/src/algorithm/sort/selectsort"
	"learnGo/src/algorithm/sort/shell"
	"learnGo/src/algorithm/sort/tool"
)

func main() {
	bubble := new(bubble.BubbleSort)
	bubble.Value, _ = tool.GenerateRandomList()
	tool.SortTool(bubble)

	sel := new(selectsort.SelectSort)
	sel.Value, _ = tool.GenerateRandomList()
	tool.SortTool(sel)

	insert := new(insertionsort.InsertionSort)
	insert.Value, _ = tool.GenerateRandomList()
	tool.SortTool(insert)

	shell := new(shellsort.ShellSort)
	shell.Value, _ = tool.GenerateRandomList()
	tool.SortTool(shell)

	merge := new(mergesort.MergeSort)
	merge.Value, _ = tool.GenerateRandomList()
	tool.SortTool(merge)

	quicksort := new(quicksort.QuickSort)
	quicksort.Value, _ = tool.GenerateRandomList()
	tool.SortTool(quicksort)

	heapsort := new(heapsort.HeapSort)
	heapsort.Value, _ = tool.GenerateRandomList()
	tool.SortTool(heapsort)
}
