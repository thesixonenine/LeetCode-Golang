package remove_duplicates_from_sorted_array

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	// Input array
	nums := []int{1, 1, 2}
	// The expected answer with correct length
	expectedNums := [...]int{1, 2}
	// Calls your implementation
	k := removeDuplicates(nums)
	if k != len(expectedNums) {
		t.Errorf("got '%v' want '%v'", k, len(expectedNums))
	}
	for i := 0; i < k; i++ {
		if nums[i] != expectedNums[i] {
			t.Errorf("got '%v' want '%v'", nums[i], expectedNums[i])
		}
	}
}
