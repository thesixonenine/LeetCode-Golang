package longest_substring_without_repeating_characters

import (
	"reflect"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		args string
		want int
	}{
		{name: "case1", args: "abcabcbb", want: 3},
		{name: "case2", args: "bbbbb", want: 1},
		{name: "case3", args: "pwwkew", want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
