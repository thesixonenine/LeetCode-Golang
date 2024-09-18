package merge_sorted_array

// merge https://leetcode.com/problems/merge-sorted-array
// nums1.length == m + n
// nums2.length == n
// 0 <= m, n <= 200
// 1 <= m + n <= 200
// -10^9 <= nums1[i], nums2[j] <= 10^9
func merge(nums1 []int, m int, nums2 []int, n int) {
	for n > 0 {
		if m > 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
func merge3(nums1 []int, m int, nums2 []int, n int) {
	for m+n > 0 {
		if n == 0 {
			break
		}
		if m > 0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[m+n-1] = nums1[m-1]
				m--
			} else {
				nums1[m+n-1] = nums2[n-1]
				n--
			}
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}
}
func merge2(nums1 []int, m int, nums2 []int, n int) {
	a := m + n - 1
	for ; a >= 0; a-- {
		if n == 0 {
			break
		}
		if m > 0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[a] = nums1[m-1]
				m--
			} else {
				nums1[a] = nums2[n-1]
				n--
			}
		} else {
			nums1[a] = nums2[n-1]
			n--
		}
	}
}
func merge1(nums1 []int, m int, nums2 []int, n int) {
	a := m + n - 1
	for ; a >= 0; a-- {
		fromM := false
		if m > 0 && n > 0 {
			// 比较后拿
			fromM = nums1[m-1] > nums2[n-1]
		} else if m > 0 {
			// 则从 m 中拿
			fromM = true
		} else if n > 0 {
			// 则从 n 中拿
		} else {
			// 都拿完了
			break
		}

		if fromM {
			nums1[a] = nums1[m-1]
			m--
		} else {
			nums1[a] = nums2[n-1]
			n--
		}
	}
}
