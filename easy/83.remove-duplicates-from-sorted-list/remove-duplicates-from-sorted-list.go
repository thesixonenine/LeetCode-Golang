package remove_duplicates_from_sorted_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := head
	for head.Next != nil {
		next := head.Next
		if head.Val == next.Val {
			head.Next = next.Next
		} else {
			head = next
		}
	}
	return ans
}
