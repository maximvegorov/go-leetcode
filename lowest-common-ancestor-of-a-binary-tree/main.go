package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p == root || q == root {
		return root
	}
	if p == q {
		return p
	}
	queue := []*pathElem{{node: root, length: 1}}
	pq := make(map[*TreeNode]*pathElem)
	for len(queue) > 0 && len(pq) < 2 {
		n := queue[0]
		queue = queue[1:]
		if n.node.Left != nil {
			ln := n.node.Left
			left := &pathElem{node: ln, prev: n, length: n.length + 1}
			queue = append(queue, left)
			if ln == p || ln == q {
				pq[ln] = left
			}
		}
		if n.node.Right != nil {
			rn := n.node.Right
			right := &pathElem{node: rn, prev: n, length: n.length + 1}
			queue = append(queue, right)
			if rn == p || rn == q {
				pq[rn] = right
			}
		}
	}
	pn := pq[p]
	qn := pq[q]
	if pn.length > qn.length {
		pn, qn = qn, pn
	}
	for qn.length > pn.length {
		qn = qn.prev
	}
	for pn.node != qn.node {
		pn = pn.prev
		qn = qn.prev
	}
	return pn.node
}

type pathElem struct {
	node   *TreeNode
	prev   *pathElem
	length int
}
