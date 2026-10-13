package convert_sorted_array_to_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	le := len(nums)
	if le == 0 {
		return nil
	}
	idx := le / 2
	root := &TreeNode{Val: nums[idx]}
	root.Left = sortedArrayToBST(nums[:idx])
	root.Right = sortedArrayToBST(nums[idx+1:])
	return root
}
