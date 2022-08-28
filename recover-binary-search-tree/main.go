package main

import "github.com/maximvegorov/go-leetcode"

type TreeNode = leetcode.TreeNode

func recoverTree(root *TreeNode) {
	w := walker{}
	w.walk(root)
	if w.q == nil {
		w.q = w.nextP
	}
	w.p.Val, w.q.Val = w.q.Val, w.p.Val
}

type walker struct {
	p     *TreeNode
	nextP *TreeNode
	q     *TreeNode
	prev  *TreeNode
}

func (w *walker) walk(node *TreeNode) {
	if w.p != nil && w.q != nil {
		return
	}
	if node.Left != nil {
		w.walk(node.Left)
	}
	if w.prev != nil {
		if w.prev.Val > node.Val {
			if w.p == nil {
				w.p = w.prev
				w.nextP = node
			} else {
				w.q = node
			}
		}
	}
	w.prev = node
	if node.Right != nil {
		w.walk(node.Right)
	}
}
