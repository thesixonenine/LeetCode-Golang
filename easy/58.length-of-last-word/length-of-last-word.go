package length_of_last_word

// https://leetcode.com/problems/length-of-last-word/
func lengthOfLastWord(s string) int {
	bytes := []byte(s)
	l := len(bytes)
	r := l - 1
	for bytes[r] == 32 {
		r--
	}
	for i := r; i >= 0; i-- {
		if bytes[i] == 32 {
			return r - i
		}
	}
	return r + 1
}
