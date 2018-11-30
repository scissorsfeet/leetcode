package _4_tree

func sortedArrayToBST(nums []int) *TreeNode {
	numsLen := len(nums)
	if numsLen == 0 {
		return nil
	}
	if numsLen == 1 {
		return &TreeNode{Val:nums[0]}
	}

	mid := numsLen / 2
	root := &TreeNode{Val:nums[mid]}
	if mid > 0 {
		root.Left = sortedArrayToBST(nums[0:mid])
	} else {
		root.Left = nil
	}
	if mid+1 < numsLen {
		root.Right = sortedArrayToBST(nums[mid+1:])
	} else {
		root.Right = nil
	}
	return root
}
