package climbing_stairs

// https://leetcode.com/problems/climbing-stairs/
func climbStairs(n int) int {
	// f(0) = 1, f(1) = 1
	// f(x) = f(x-1) + f(x-2)
	x, y, ans := 0, 0, 1
	for i := 0; i < n; i++ {
		x = y
		y = ans
		ans = x + y
	}
	return ans
}
