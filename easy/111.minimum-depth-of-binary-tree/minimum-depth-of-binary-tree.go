package minimum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return minDepth(root.Right) + 1
	}
	if root.Right == nil {
		return minDepth(root.Left) + 1
	}
	ld := minDepth(root.Left)
	rd := minDepth(root.Right)
	if ld >= rd {
		return rd + 1
	} else {
		return ld + 1
	}
}
