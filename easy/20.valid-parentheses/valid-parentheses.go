package valid_parentheses

// https://leetcode.com/problems/valid-parentheses/
func isValid(s string) bool {
	bytes := []byte(s)
	l := len(bytes)
	if l == 0 || l%2 != 0 {
		return false
	}
	// 最多支持的深度
	maxLength := l / 2
	stack := make([]byte, 0)
	typeMap := map[byte]byte{')': '(', ']': '[', '}': '{'}
	for _, b := range bytes {
		m, ok := typeMap[b]
		if !ok {
			// 需要入栈
			if len(stack) == maxLength {
				return false
			}
			stack = append(stack, b)
		} else {
			top := len(stack) - 1
			if top < 0 {
				return false
			}
			if m == stack[top] {
				stack = stack[:top]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
