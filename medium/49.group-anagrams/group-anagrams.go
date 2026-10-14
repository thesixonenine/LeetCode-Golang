package group_anagrams

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	m := make(map[string][]string, len(strs))
	for _, w := range strs {
		bs := []byte(w)
		slices.Sort(bs)
		b := string(bs)
		m[b] = append(m[b], w)
	}
	ans := make([][]string, 0)
	for _, v := range m {
		if len(v) == 0 {
			continue
		}
		ans = append(ans, v)
	}
	return ans
}
