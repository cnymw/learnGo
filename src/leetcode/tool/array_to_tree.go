package tool

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func ArrayToTree(array []int) *TreeNode {
	nodeArray := make([]TreeNode, 0)
	for _, i := range array {
		nodeArray = append(nodeArray, TreeNode{Val: i})
	}

	if len(nodeArray) > 0 {
		for i := 0; i < len(array)/2; i++ {
			if nodeArray[left(i)].Val != 0 {
				nodeArray[i].Left = &nodeArray[left(i)]
			}
			if nodeArray[right(i)].Val != 0 {
				nodeArray[i].Right = &nodeArray[right(i)]
			}
		}
	}

	return &nodeArray[0]
}

func left(v int) int {
	return 2*v + 1
}

func right(v int) int {
	return 2*v + 2
}
