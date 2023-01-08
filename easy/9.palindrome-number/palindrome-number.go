package palindrome_number

import (
	"fmt"
	"strconv"
	"syscall"
)

// https://leetcode.cn/problems/palindrome-number
func PalindromeNumber1(x int) bool {
	// -231 <= x <= 231 - 1
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	}
	// 正整数 转为 字符串
	xStr := strconv.Itoa(x)
	// 字符串 转为 字节数组
	for i, b := range xStr {
		fmt.Println(i, b)
	}
	bytes, _ := syscall.ByteSliceFromString(xStr)
	t := len(bytes)
	for i := 0; i < t; i++ {

	}
	return false
}

func PalindromeNumber2(x int) bool {
	// -231 <= x <= 231 - 1
	if x < 0 {
		return false
	} else if x == 0 {
		return true
	} else {
		// 正整数
	}
	return false
}
