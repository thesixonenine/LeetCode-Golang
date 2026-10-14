package group_anagrams

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want [][]string
	}{
		{name: "case1", args: []string{"eat", "tea", "tan", "ate", "nat", "bat"}, want: [][]string{[]string{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}}},
		{name: "case2", args: []string{""}, want: [][]string{[]string{""}}},
		{name: "case3", args: []string{"a"}, want: [][]string{[]string{"a"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
