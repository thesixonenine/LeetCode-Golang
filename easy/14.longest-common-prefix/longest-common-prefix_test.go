package longest_common_prefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	str := []string{"flower", "flow", "flight"}
	want := "fl"

	got := longestCommonPrefix(str)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}

	str = []string{"dog", "racecar", "car"}
	want = ""

	got = longestCommonPrefix(str)
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
