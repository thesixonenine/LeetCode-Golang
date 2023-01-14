package remove_duplicates_from_sorted_array

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func removeDuplicates(nums []int) int {
	l := len(nums)
	ret := 0
	for i := 1; i < l; i++ {
		if nums[ret] != nums[i] {
			ret = ret + 1
			nums[ret] = nums[i]
		}
	}
	return ret + 1
}
