package search_in_a_binary_search_tree

import (
	"github.com/maximvegorov/go-leetcode"
)

type TreeNode = leetcode.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if root.Val == val {
			return root
		} else if root.Val > val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
