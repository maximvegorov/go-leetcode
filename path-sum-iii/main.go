package path_sum_iii

import "github.com/maximvegorov/go-leetcode"

type TreeNode = leetcode.TreeNode

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	c := pathSumCounter{target: targetSum}
	c.walk(&pathElem{n: root})
	return c.count
}

type pathSumCounter struct {
	target int
	count  int
}

func (c *pathSumCounter) walk(tail *pathElem) {
	sum := c.target
	for p := tail; p != nil; p = p.next {
		sum -= p.n.Val
		if sum == 0 {
			c.count++
		}
	}

	if left := tail.n.Left; left != nil {
		c.walk(&pathElem{n: left, next: tail})
	}

	if right := tail.n.Right; right != nil {
		c.walk(&pathElem{n: right, next: tail})
	}
}

type pathElem struct {
	n    *TreeNode
	next *pathElem
}
