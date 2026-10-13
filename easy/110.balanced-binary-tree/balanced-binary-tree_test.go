package balanced_binary_tree

import (
	"reflect"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	tests := []struct {
		name string
		args *TreeNode
		want bool
	}{
		{name: "case1", args: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, want: true},
		{name: "case2", args: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 2}}, want: false},
		{name: "case3", args: nil, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}
