package main

import "github.com/maximvegorov/go-leetcode"

type TreeNode = leetcode.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	inorderMap := make(map[int]int, len(preorder))
	for i, v := range inorder {
		inorderMap[v] = i
	}
	b := builder{preorder: preorder, inorder: inorderMap}
	return b.build(0, 0, len(inorder))
}

type builder struct {
	preorder []int
	inorder  map[int]int
}

func (b *builder) build(rootIndex int, inorderLeft, inorderRight int) *TreeNode {
	val := b.preorder[rootIndex]
	root := &TreeNode{Val: val}
	inorderPos := b.inorder[val]
	leftCount := inorderPos - inorderLeft
	rightCount := inorderRight - inorderPos - 1
	if leftCount > 0 {
		root.Left = b.build(rootIndex+1, inorderLeft, inorderPos)
	}
	if rightCount > 0 {
		root.Right = b.build(rootIndex+leftCount+1, inorderPos+1, inorderRight)
	}
	return root
}
