package merge_two_sorted_lists

import "testing"

func TestIsValid(t *testing.T) {
	testAllList := make([][3]ListNode, 0)

	testList1 := [3]ListNode{}
	node1 := ListNode{Val: 4}
	node2 := ListNode{Val: 2, Next: &node1}
	// 输入链表: [1,2,4]
	testList1[0] = ListNode{Val: 1, Next: &node2}

	node3 := ListNode{Val: 4}
	node4 := ListNode{Val: 3, Next: &node3}
	// 输入链表: [1,3,4]
	testList1[1] = ListNode{Val: 1, Next: &node4}

	node5 := ListNode{Val: 4}
	node6 := ListNode{Val: 4, Next: &node5}
	node7 := ListNode{Val: 3, Next: &node6}
	node8 := ListNode{Val: 2, Next: &node7}
	node9 := ListNode{Val: 1, Next: &node8}
	// 输出链表: [1,1,2,3,4,4]
	testList1[2] = ListNode{Val: 1, Next: &node9}
	// 将测试用例写入AllList
	testAllList = append(testAllList, testList1)

	for _, item := range testAllList {
		got := mergeTwoLists(&item[0], &item[1])
		want := &item[2]
		if want.Next == nil {
			if want.Val == got.Val {
				continue
			} else {
				t.Errorf("got '%v' want '%v'", got.Val, want.Val)
			}
		}
		for want.Next != nil {
			if got == nil {
				t.Errorf("got '%v' want '%v'", got, want)
			}
			if want.Val == got.Val {
				want = want.Next
				got = got.Next
			} else {
				t.Errorf("got '%v' want '%v'", got.Val, want.Val)
			}
		}
	}

}
