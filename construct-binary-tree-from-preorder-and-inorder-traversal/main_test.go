package main

import (
	"testing"
)

func Test_buildTree(t *testing.T) {
	type args struct {
		preorder []int
		inorder  []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{"1", args{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buildTree(tt.args.preorder, tt.args.inorder)
		})
	}
}
