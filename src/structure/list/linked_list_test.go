package list

import (
	"fmt"
	"testing"
)

func TestList_ShowList(t *testing.T) {
	list := List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.ShowList()
}

func TestList_Contain(t *testing.T) {
	list := List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	fmt.Println(list.Contain(2))
	fmt.Println(list.Contain(5))
}

func TestList_IsEmpty(t *testing.T) {
	list := List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	fmt.Println(list.IsEmpty())
}

func TestList_Length(t *testing.T) {
	list := List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	fmt.Println(list.Length())
}

func TestList_Remove(t *testing.T) {
	list := List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Remove(2)
	list.ShowList()
}

func TestList_RemoveAtIndex(t *testing.T) {
	list := List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.RemoveAtIndex(2)
	list.ShowList()
}

func TestList_Insert(t *testing.T) {
	list := List{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	list.Insert(2, 5)
	list.ShowList()
}
