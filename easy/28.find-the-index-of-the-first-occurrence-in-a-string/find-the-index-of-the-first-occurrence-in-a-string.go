package find_the_index_of_the_first_occurrence_in_a_string

import "strings"

// strStr https://leetcode.com/problems/find-the-index-of-the-first-occurrence-in-a-string/
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
