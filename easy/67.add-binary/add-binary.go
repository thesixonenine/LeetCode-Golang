package add_binary

import "strconv"

// https://leetcode.com/problems/add-binary/
func addBinary(a string, b string) string {
	aByte := []byte(a)
	bByte := []byte(b)
	aLen := len(aByte) - 1
	bLen := len(bByte) - 1

	if aLen == 0 && string(aByte[0]) == "0" && bLen == 0 && string(bByte[0]) == "0" {
		return "0"
	}

	plus := 0
	ans := ""
	for {
		x := 0
		y := 0
		if aLen >= 0 {
			x, _ = strconv.Atoi(string(aByte[aLen]))
		}
		if bLen >= 0 {
			y, _ = strconv.Atoi(string(bByte[bLen]))
		}

		sum := x + y + plus

		if aLen <= 0 && bLen <= 0 && sum == 0 {
			break
		}
		plus = sum / 2
		ans = strconv.Itoa(sum%2) + ans
		aLen--
		bLen--

	}

	return ans
}
