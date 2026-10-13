package minimum_depth_of_binary_tree

import (
	"reflect"
	"testing"
)

func Test_minDepth(t *testing.T) {
	tests := []struct {
		name string
		args *TreeNode
		want int
	}{
		{name: "case1", args: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, want: 2},
		{name: "case2", args: &TreeNode{Val: 2, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 6}}}}}, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
