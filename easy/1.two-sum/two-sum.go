package two_sum

// https://leetcode.com/problems/two-sumcd
func twoSum(numbs []int, target int) []int {
	// 使用map
	intMap := map[int]int{}
	for k, value := range numbs {
		v, ok := intMap[target-value]
		if ok {
			return []int{v, k}
		} else {
			intMap[value] = k
		}
	}
	return nil
}
