package balanced_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	c := height(root.Left) - height(root.Right)
	if c < 0 {
		c = -c
	}
	if c > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	lh := height(node.Left)
	rh := height(node.Right)
	if lh >= rh {
		return lh + 1
	} else {
		return rh + 1
	}
}
