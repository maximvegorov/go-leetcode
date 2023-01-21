package flattenbinarytreetolinkedlist

import (
	"github.com/maximvegorov/go-leetcode"
)

type TreeNode = leetcode.TreeNode

func flatten(root *TreeNode) {
	flattenToList(root)
}

func flattenToList(root *TreeNode) (first *TreeNode, last *TreeNode) {
	if root == nil {
		return nil, nil
	}
	lf, ll := flattenToList(root.Left)
	rf, rl := flattenToList(root.Right)
	root.Left = nil
	if lf != nil {
		root.Right = lf
	} else {
		ll = root
	}
	if rf != nil {
		ll.Right = rf
	} else {
		rl = ll
	}
	rl.Right = nil
	return root, rl
}
