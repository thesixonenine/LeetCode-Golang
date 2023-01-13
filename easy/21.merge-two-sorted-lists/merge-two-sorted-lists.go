package merge_two_sorted_lists

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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val > list2.Val {
		// 以list2开始
		temp := list1
		list1 = list2
		list2 = temp
	}
	if list1.Next == nil {
		list1.Next = list2
		return list1
	}
	if list2.Next == nil {
		list2.Next = list1.Next
		list1.Next = list2
		return list1
	}
	for {
		if list1.Next != nil {
			if list1.Val <= list2.Val && list2.Val <= list1.Next.Val {
				list1.Next = &ListNode{Val: list2.Val, Next: list1.Next}
			}
		}

		if list1.Next != nil && list2.Next != nil {
			break
		}
	}

	node5 := ListNode{Val: 4}
	node6 := ListNode{Val: 4, Next: &node5}
	node7 := ListNode{Val: 3, Next: &node6}
	node8 := ListNode{Val: 2, Next: &node7}
	node9 := ListNode{Val: 1, Next: &node8}
	return &ListNode{Val: 1, Next: &node9}
}
