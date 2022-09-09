package most_stones_removed_with_same_row_or_column

func removeStones(stones [][]int) int {
	if len(stones) == 0 {
		return 0
	}
	fu := NewFindUnion(len(stones))
	for i, p1 := range stones {
		xi := p1[0]
		yi := p1[1]
		for j := i + 1; j < len(stones); j++ {
			p2 := stones[j]
			xj := p2[0]
			yj := p2[1]
			if xi == xj || yi == yj {
				fu.Union(i, j)
			}
		}
	}
	return len(stones) - fu.Count()
}

type FindUnion struct {
	forest []FindUnionNode
}

func NewFindUnion(size int) *FindUnion {
	forest := make([]FindUnionNode, size)
	for t := range forest {
		n := &forest[t]
		n.parent = n
		n.size = 1
	}
	return &FindUnion{forest: forest}
}

type FindUnionNode struct {
	parent *FindUnionNode
	size   int
}

func (fu *FindUnion) Find(i int) *FindUnionNode {
	n := &fu.forest[i]
	for n.parent != n {
		n = n.parent
	}
	return n
}

func (fu *FindUnion) Union(i, j int) {
	ni := fu.Find(i)
	nj := fu.Find(j)
	if ni == nj {
		return
	}
	if ni.size < nj.size {
		ni, nj = nj, ni
	}
	ni.size += nj.size
	nj.parent = ni
}

func (fu *FindUnion) Count() int {
	res := 0
	for t := range fu.forest {
		n := &fu.forest[t]
		if n == n.parent {
			res++
		}
	}
	return res
}
