package main

import "github.com/maximvegorov/go-leetcode"

type TreeNode = leetcode.TreeNode

func kthSmallest(root *TreeNode, k int) int {
	cur := 0
	n := kthElement(root, &cur, k)
	if n == nil {
		return -1
	} else {
		return *n
	}
}

func kthElement(root *TreeNode, cur *int, k int) *int {
	if root == nil {
		return nil
	}
	if root.Left != nil {
		if n := kthElement(root.Left, cur, k); n != nil {
			return n
		}
	}
	*cur++
	if *cur == k {
		return &root.Val
	}
	if root.Right != nil {
		if n := kthElement(root.Right, cur, k); n != nil {
			return n
		}
	}
	return nil
}
