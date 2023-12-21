package constructquadtree

import "github.com/maximvegorov/go-leetcode"

type Node = leetcode.QNode

func construct(grid [][]int) *Node {
	return constructNode(grid, 0, 0, len(grid))
}

func constructNode(grid [][]int, x, y int, sz int) *Node {
	if sz == 1 {
		return &Node{Val: grid[y][x] == 1, IsLeaf: true}
	}

	sz /= 2

	tl := constructNode(grid, x, y, sz)
	tr := constructNode(grid, x+sz, y, sz)
	bl := constructNode(grid, x, y+sz, sz)
	br := constructNode(grid, x+sz, y+sz, sz)

	isLeaf := tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf
	if !isLeaf {
		return &Node{TopLeft: tl, TopRight: tr, BottomLeft: bl, BottomRight: br}
	}

	if !(tl.Val || tr.Val || bl.Val || br.Val) {
		return &Node{IsLeaf: true}
	} else if tl.Val && tr.Val && bl.Val && br.Val {
		return &Node{Val: true, IsLeaf: true}
	} else {
		return &Node{TopLeft: tl, TopRight: tr, BottomLeft: bl, BottomRight: br}
	}
}
