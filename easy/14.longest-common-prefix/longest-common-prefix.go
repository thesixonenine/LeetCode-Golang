package longest_common_prefix

func longestCommonPrefix(str []string) string {
	// 所有字符串横向扫描, 从0到最短的str的length
	strLength := len(str)
	if strLength == 1 {
		return str[0]
	}
	byteMap := map[int][]byte{}
	minLength := 200
	for i, elem := range str {
		byteMap[i] = []byte(elem)
		l := len(elem)
		if l < minLength {
			minLength = l
		}
	}
	if minLength == 0 {
		return ""
	}
	startByteArr := byteMap[0]
	onceDiff := false
	for i := 0; i < minLength; i++ {
		// 判断所有字符串的第i个字符是否相等
		firstByte := startByteArr[i]
		hasDiff := false
		for j := 1; j < strLength; j++ {
			o := byteMap[j][i]
			if firstByte != o {
				// 有不等的
				hasDiff = true
				onceDiff = true
				break
			}
		}
		if minLength == 1 {
			if !hasDiff {
				return str[0][0:1]
			} else {
				return ""
			}
		}
		if hasDiff {
			return str[0][0:i]
		}
	}
	if onceDiff {
		return ""
	}
	return str[0][0:minLength]
}

func longestCommonPrefix1(str []string) string {
	length := len(str)
	if length == 0 {
		return ""
	}
	if length == 1 {
		return str[0]
	}
	item1Length := len(str[0])
	if item1Length == 0 {
		return ""
	}
	// 两个for循环
	// 横向循环第一个字符串
	for i := 0; i < item1Length; i++ {
		// 纵向循环
		for j := 1; j < length; j++ {
			if str[j][i] != str[j-1][i] {
				return str[0][:i]
			}
		}
	}
	return ""
}
