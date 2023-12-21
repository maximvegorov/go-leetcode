package wordsearchii

func findWords(board [][]byte, words []string) []string {
	trie := buildTrie(words)
	q := make([]*queueItem, 0, 2)
	for i, row := range board {
		for j, c := range row {
			tn := trie.children[c]
			if tn == nil {
				continue
			}
			q = append(q, &queueItem{r: byte(i), c: byte(j), n: tn})
		}
	}
	dict := make(map[string]struct{})
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.n.word != nil {
			dict[*cur.n.word] = struct{}{}
		}
		if r := int(cur.r) - 1; r >= 0 {
			n := cur.n.children[board[r][cur.c]]
			if n != nil {
				q = append(q, &queueItem{r: byte(r), c: cur.c, n: n})
			}
		}
		if c := int(cur.c) - 1; c >= 0 {
			n := cur.n.children[board[cur.r][c]]
			if n != nil {
				q = append(q, &queueItem{r: cur.r, c: byte(c), n: n})
			}
		}
		if r := cur.r + 1; int(r) < len(board) {
			n := cur.n.children[board[r][cur.c]]
			if n != nil {
				q = append(q, &queueItem{r: r, c: cur.c, n: n})
			}
		}
		if c := cur.c + 1; int(c) < len(board[cur.r]) {
			n := cur.n.children[board[cur.r][c]]
			if n != nil {
				q = append(q, &queueItem{r: cur.r, c: c, n: n})
			}
		}
	}
	res := make([]string, len(dict))
	for w := range dict {
		res = append(res, w)
	}
	return res
}

type queueItem struct {
	r, c byte
	n    *trieNode
}

func buildTrie(words []string) *trieNode {
	root := &trieNode{}
	for _, w := range words {
		p := root
		n := len(w) - 1
		for i := 0; i < n; i++ {
			p = p.add(w[i], nil)
		}
		p.add(w[n], &w)
	}
	return root
}

type trieNode struct {
	children map[byte]*trieNode
	word     *string
}

func (n *trieNode) add(c byte, word *string) *trieNode {
	child := n.children[c]
	if child == nil {
		child = new(trieNode)
		if n.children == nil {
			n.children = make(map[byte]*trieNode)
		}
		n.children[c] = child
	}
	child.word = word
	return child
}
