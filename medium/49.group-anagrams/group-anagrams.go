package group_anagrams

import (
	"maps"
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
	return slices.Collect(maps.Values(m))
}
