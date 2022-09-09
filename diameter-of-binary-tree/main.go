package diameter_of_binary_tree

import "github.com/maximvegorov/go-leetcode"

type TreeNode = leetcode.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	w := walker{}
	w.walk(root)
	return w.d
}

type walker struct {
	d int
}

func (w *walker) walk(root *TreeNode) int {
	leftHeight := 0
	if root.Left != nil {
		leftHeight = w.walk(root.Left)
	}
	rightHeight := 0
	if root.Right != nil {
		rightHeight = w.walk(root.Right)
	}
	d := leftHeight + rightHeight
	if d > w.d {
		w.d = d
	}
	h := leftHeight
	if rightHeight > h {
		h = rightHeight
	}
	return h + 1
}
