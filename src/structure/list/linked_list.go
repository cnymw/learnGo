package list

import "fmt"

type Object interface{}

type Node struct {
	Data Object //定义数据域
	Next *Node  //定义地址域(指向下一个表的地址)
}

type List struct {
	headNode *Node //头节点
}

// 判断链表是否为空
func (list *List) IsEmpty() bool {
	return list.headNode == nil
}

// 获取列表长度
func (list *List) Length() int {
	//获取链表头结点
	cur := list.headNode
	//定义一个计数器，初始值为0
	count := 0

	for cur != nil {
		//如果头节点不为空，则count++
		count++
		//对地址进行逐个位移
		cur = cur.Next
	}
	return count
}

// 从链表头部添加元素
func (list *List) Add(data Object) *Node {
	node := &Node{Data: data}
	node.Next = list.headNode
	list.headNode = node
	return node
}

// 从链表尾部添加元素
func (list *List) Append(data Object) {
	node := &Node{Data: data}
	if list.IsEmpty() {
		list.headNode = node
	} else {
		cur := list.headNode
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = node
	}
}

// 在链表指定位置添加元素
func (list *List) Insert(index int, data Object) {
	if index < 0 {
		list.Add(data)
	} else if index > list.Length() {
		list.Append(data)
	} else {
		pre := list.headNode
		count := 0
		for count < index-1 {
			pre = pre.Next
			count++
		}
		node := &Node{Data: data}
		node.Next = pre.Next
		pre.Next = node
	}
}

// 删除链表指定值的元素
func (list *List) Remove(data Object) {
	pre := list.headNode
	if pre.Data == data {
		list.headNode = pre.Next
	} else {
		for pre.Next != nil {
			if pre.Next.Data == data {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
	}
}

// 删除指定位置的元素
func (list *List) RemoveAtIndex(index int) {
	pre := list.headNode
	if index < 0 {
		list.headNode = pre.Next
	} else if index > list.Length() {
		return
	} else {
		count := 0
		for count != index-1 && pre.Next != nil {
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}

// 查看链表是否包含某个元素
func (list *List) Contain(data Object) bool {
	cur := list.headNode
	for cur != nil {
		if cur.Data == data {
			return true
		}
		cur = cur.Next
	}
	return false
}

// 遍历链表所有节点
func (list *List) ShowList() {
	if !list.IsEmpty() {
		cur := list.headNode
		for cur != nil {
			fmt.Println(cur.Data)
			cur = cur.Next
		}
	}
}
