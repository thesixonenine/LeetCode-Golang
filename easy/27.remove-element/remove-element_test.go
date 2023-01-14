package remove_element

import "testing"

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	want := 2
	wantArr := []int{2, 2}
	got := removeElement(nums, val)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
	for i, item := range wantArr {
		if nums[i] != item {
			t.Errorf("got '%v' want '%v'", nums[i], item)
		}
	}
}
