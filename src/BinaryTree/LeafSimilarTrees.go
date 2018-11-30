package BinaryTree

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := make([]int, 0, 1024)
	leafs2 := make([]int, 0, 1024)
	getLeafs(root1, &leafs1)
	getLeafs(root2, &leafs2)

	var res bool = true
	if len(leafs1) != len(leafs2) {
		res = false
	} else {
		for i:=0;i<len(leafs1);i++ {
			if leafs1[i] != leafs2[i] {
				res = false
				break
			}
		}
	}

	return res
}

func getLeafs(root *TreeNode, leafs *[]int) {
	if nil == root {
		return
	}
	if nil == root.Left && nil == root.Right {
		*leafs = append(*leafs, root.Val)
		return
	}

	if nil != root.Left {
		getLeafs(root.Left, leafs)
	}
	if nil != root.Right {
		getLeafs(root.Right, leafs)
	}
}