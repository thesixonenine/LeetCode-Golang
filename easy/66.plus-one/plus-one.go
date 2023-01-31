package plus_one

// https://leetcode.com/problems/plus-one/
func plusOne(digits []int) []int {
	l := len(digits)
	hasPlus := digits[l-1] == 9
	leftPushOne := false
	for i := l - 1; i >= 0; i-- {
		if hasPlus {
			if digits[i] == 9 {
				digits[i] = 0
				leftPushOne = i == 0
			} else {
				digits[i] = digits[i] + 1
				break
			}
		} else {
			digits[i] = digits[i] + 1
			break
		}
	}
	var ans []int
	if leftPushOne {
		ans = make([]int, l+1)
		ans[0] = 1
		for i := 1; i < l; i++ {
			ans[i] = digits[i-1]
		}

	} else {
		ans = make([]int, l)
		for i := 0; i < l; i++ {
			ans[i] = digits[i]
		}
	}
	return ans
}
