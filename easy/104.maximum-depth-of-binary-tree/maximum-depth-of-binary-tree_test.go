package maximum_depth_of_binary_tree

import "testing"

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "case1", args: args{root: &TreeNode{Val: 3,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}}, want: 3},
		{name: "case2", args: args{root: &TreeNode{Val: 1,
			Right: &TreeNode{Val: 2}}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
