package remove_duplicates_from_sorted_list

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 1}
	l3 := &ListNode{Val: 2}
	l1.Next = l2
	l2.Next = l3

	want := &ListNode{Val: 1}
	w1 := &ListNode{Val: 2}
	want.Next = w1

	got := deleteDuplicates(l1)

	for want != nil {
		if want.Val != got.Val {
			t.Errorf("got '%v' want '%v'", got.Val, want.Val)
		}
		want = want.Next
		got = got.Next
	}
}
