package path_sum

import (
	"reflect"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	tests := []struct {
		name      string
		args      *TreeNode
		targetSum int
		want      bool
	}{
		{name: "case1", args: &TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 11, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 13}, Right: &TreeNode{Val: 4, Right: &TreeNode{Val: 1}}}}, targetSum: 22, want: true},
		{name: "case2", args: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, targetSum: 5, want: false},
		{name: "case3", args: nil, targetSum: 0, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args, tt.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
