package palindrome_number

import (
	"strconv"
)

// https://leetcode.com/problems/palindrome-number
func isPalindrome(x int) bool {
	// -231 <= x <= 231 - 1
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	// 正整数 转为 字符串
	xStr := strconv.Itoa(x)
	// 字符串 转为 字符数组
	bytes := []byte(xStr)
	t := len(bytes)
	for i := 0; i <= t/2; i++ {
		head := bytes[i]
		tail := bytes[t-i-1]
		if head != tail {
			return false
		}
	}
	return true
}

func isPalindrome1(x int) bool {
	// -231 <= x <= 231 - 1
	if x < 0 {
		return false
	} else if x <= 9 {
		return true
	} else if x%10 == 0 {
		return false
	}
	// 翻转一半的数字
	var half int
	for {
		if x < 10 {
			break
		}
		half = half*10 + x%10
		if x < half {
			break
		} else if x == half {
			return true
		}
		x = x / 10
		if x < half {
			break
		} else if x == half {
			return true
		}
	}
	return half == x
}
