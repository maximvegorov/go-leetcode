package clonegraph

import (
	"slices"

	"github.com/maximvegorov/go-leetcode"
)

type Node = leetcode.GNode

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[*Node]*Node)
	q := []*Node{node}
	var processed []*Node
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if _, ok := visited[cur]; ok {
			continue
		}
		clone := &Node{Val: cur.Val, Neighbors: slices.Clone(cur.Neighbors)}
		visited[cur] = clone
		processed = append(processed, clone)
		for _, n := range clone.Neighbors {
			if _, ok := visited[n]; ok {
				continue
			}
			q = append(q, n)
		}
	}
	for _, n := range processed {
		for i := 0; i < len(n.Neighbors); i++ {
			n.Neighbors[i] = visited[n.Neighbors[i]]
		}
	}
	return visited[node]
}
