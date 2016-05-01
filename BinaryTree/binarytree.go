package BinaryTree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//新建叶子节点
func NewTreeNode(val int, left *TreeNode, right *TreeNode) *TreeNode {
	return &TreeNode{val, left, right}
}

func CreateByArray(nodes []int) (root *TreeNode) {
	return createByArray(nodes, 1, len(nodes))
}

//通过数组建立
func createByArray(nodes []int, rootIndex int, nodesLen int) (root *TreeNode) {
	if rootIndex <= nodesLen {
		if nodes[rootIndex-1] >= 0 {
			var left, right *TreeNode
			if rootIndex*2 <= nodesLen {
				if nodes[rootIndex*2-1] >= 0 {
					left = createByArray(nodes, rootIndex*2, nodesLen)
				}
			}
			if rootIndex*2+1 <= nodesLen {
				if nodes[rootIndex*2] >= 0 {
					right = createByArray(nodes, rootIndex*2+1, nodesLen)
				}
			}
			root = NewTreeNode(nodes[rootIndex-1], left, right)
		}
	}
	return root
}

func (root *TreeNode) MaxDepth() int {
	if root == nil {
		return 0
	}
	leftDepth := root.Left.MaxDepth()
	rightDepth := root.Right.MaxDepth()
	if leftDepth >= rightDepth {
		return leftDepth + 1
	} else {
		return rightDepth + 1
	}
}

func (root *TreeNode) LevelTraverse() {
	if root == nil {
		fmt.Println("empty tree")
		return
	}
	nodeQueue := make(chan *TreeNode, 10)
	nodeQueue <- root
	i := 0
	for len(nodeQueue) > 0 {
		tmpNode := <-nodeQueue
		if tmpNode != nil {
			fmt.Printf("%d ", tmpNode.Val)
			nodeQueue <- tmpNode.Left
			nodeQueue <- tmpNode.Right
		} else {
			fmt.Printf("# ")
		}
		i++
	}
	fmt.Println()
}
