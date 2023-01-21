package flattenbinarytreetolinkedlist

import "testing"

func Test_flatten(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"1",
			args{
				root: &TreeNode{
					Val:   1,
					Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
					Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flatten(tt.args.root)
		})
	}
}
