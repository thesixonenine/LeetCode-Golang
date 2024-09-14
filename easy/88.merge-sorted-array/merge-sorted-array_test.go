package merge_sorted_array

import "testing"

func TestMerge(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	nums1want := []int{1, 2, 2, 3, 5, 6}

	merge(nums1, m, nums2, n)

	for i, want := range nums1want {
		if nums1[i] != want {
			t.Errorf("got '%d' want '%d'", nums1[i], want)
		}
	}
}
