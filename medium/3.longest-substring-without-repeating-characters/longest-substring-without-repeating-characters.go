package longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	ans := 0
	left := 0
	lastIndex := make(map[byte]int)
	for j := 0; j < len(s); j++ {
		if index, ok := lastIndex[s[j]]; ok {
			if index > left {
				left = index
			}
		}

		if j-left+1 > ans {
			ans = j - left + 1
		}

		lastIndex[s[j]] = j + 1
	}
	return ans
}
