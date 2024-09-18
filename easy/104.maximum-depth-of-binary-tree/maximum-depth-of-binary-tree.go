package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ld := maxDepth(root.Left) + 1
	rd := maxDepth(root.Right) + 1
	if ld > rd {
		return ld
	}
	return rd
}
