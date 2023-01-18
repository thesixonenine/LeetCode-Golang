package search_insert_position

import "testing"

func TestSearchInsert(t *testing.T) {
	nums := []int{1, 3, 5, 6}
	target := 5
	want := 2
	got := searchInsert(nums, target)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}

	nums = []int{1, 3, 5, 6}
	target = 2
	want = 1
	got = searchInsert(nums, target)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}

	nums = []int{1, 3, 5, 6}
	target = 7
	want = 4
	got = searchInsert(nums, target)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
