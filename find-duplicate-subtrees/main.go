package find_duplicate_subtrees

import "github.com/maximvegorov/go-leetcode"

type TreeNode = leetcode.TreeNode

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	w := treeWalker{}
	w.walk(root)
	f := newDuplicateFinder()
	for _, n := range w.nodes {
		f.add(n)
	}
	return f.duplicates()
}

type sameStatus byte

const (
	UNKNOWN sameStatus = iota
	SAME
	NOT_SAME
)

type duplicateFinder struct {
	sameCache map[nodePair]sameStatus
	unique    []*TreeNode
	freqs     map[*TreeNode]int
}

type nodePair struct {
	node1 *TreeNode
	node2 *TreeNode
}

func newDuplicateFinder() *duplicateFinder {
	return &duplicateFinder{
		sameCache: make(map[nodePair]sameStatus),
		freqs:     make(map[*TreeNode]int),
	}
}

func (f *duplicateFinder) add(root *TreeNode) {
	for _, n := range f.unique {
		if f.isSame(n, root) {
			f.freqs[n]++
			return
		}
	}
	f.unique = append(f.unique, root)
	f.freqs[root] = 1
}

func (f *duplicateFinder) isSame(root1 *TreeNode, root2 *TreeNode) bool {
	pair := nodePair{node1: root1, node2: root2}
	if s := f.sameCache[pair]; s != UNKNOWN {
		return s == SAME
	}
	s := f.checkIsSame(root1, root2)
	f.sameCache[pair] = s
	return s == SAME
}

func (f *duplicateFinder) checkIsSame(root1 *TreeNode, root2 *TreeNode) sameStatus {
	if root1 == nil {
		if root2 == nil {
			return SAME
		} else {
			return NOT_SAME
		}
	} else if root2 == nil {
		if root1 == nil {
			return SAME
		} else {
			return NOT_SAME
		}
	}
	if root1.Val != root2.Val {
		return NOT_SAME
	}
	if !f.isSame(root1.Left, root2.Left) {
		return NOT_SAME
	} else if !f.isSame(root1.Right, root2.Right) {
		return NOT_SAME
	} else {
		return SAME
	}
}

func (f *duplicateFinder) duplicates() []*TreeNode {
	var res []*TreeNode
	for k, v := range f.freqs {
		if v == 1 {
			continue
		}
		res = append(res, k)
	}
	return res
}

type treeWalker struct {
	nodes []*TreeNode
}

func (w *treeWalker) walk(root *TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		w.walk(root.Left)
	}
	if root.Right != nil {
		w.walk(root.Right)
	}
	w.nodes = append(w.nodes, root)
}
