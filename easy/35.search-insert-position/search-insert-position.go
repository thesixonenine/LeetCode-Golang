package search_insert_position

// https://leetcode.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	// 在一个有序数组中找第一个大于等于 target 的元素的下标
	ans := len(nums)
	r := ans - 1
	l := 0
	var mid int
	for l <= r {
		mid = (l + r) / 2
		if nums[mid] >= target {
			ans = mid
			r = mid - 1
		} else {
			// 查右边
			l = mid + 1
		}
	}
	return ans
}
