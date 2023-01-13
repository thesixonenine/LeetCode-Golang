package roman_to_integer

// https://leetcode.com/problems/roman-to-integer/
func romanToInt(s string) int {
	byteValueMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	bytes := []byte(s)
	l := len(bytes)
	r := 0
	for i, b := range bytes {
		v := byteValueMap[b]
		if i+1 < l {
			// 不是最后一个数字
			next := bytes[i+1]
			if b == 'I' && (next == 'V' || next == 'X') {
				v = 0 - v
			}
			if b == 'X' && (next == 'L' || next == 'C') {
				v = 0 - v
			}
			if b == 'C' && (next == 'D' || next == 'M') {
				v = 0 - v
			}
		}
		r = r + v
	}
	return r
}
