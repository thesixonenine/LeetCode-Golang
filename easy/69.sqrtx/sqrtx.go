package sqrtx

import "math"

// https://leetcode.com/problems/sqrtx/
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	// ans^2 - C = 0
	// 取 ans = C 的点来找切线
	// 求导得 y = 2ans 即斜率是2ans, 且过点(anx, ans^2 - C), 则该切线的解析式为
	// y = 2•ans•(x - ans) + ans^2 - C
	// 该切线的零点即是下一个ans
	// ans2 = 0.5 * (ans + C/ans)
	// 而当 两个ans 之间非常接近时(两者的差值小于一个极小数e), 则可以认为是答案

	e := 1e-7
	var x0 = float64(x)
	var ans = x0
	for {
		ans = 0.5 * (x0 + float64(x)/x0)
		if math.Abs(ans-x0) < e {
			break
		}
		x0 = ans
	}
	return int(x0)
}
