package pacific_atlantic_waterflow

func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	if m == 0 {
		return [][]int{}
	}
	n := len(heights[0])

	preachable := calcPacificReachable(m, n, heights)
	areachable := calcAtlanticReachable(m, n, heights)

	var res [][]int
	for i := range heights {
		for j := range heights[i] {
			c := cell{r: byte(i), c: byte(j)}
			if preachable[c] && areachable[c] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func calcPacificReachable(m, n int, heights [][]int) map[cell]bool {
	res := make(map[cell]bool)

	q := make([]cell, 0, m+n)
	visited := make(map[cell]bool)
	for j := range heights[0] {
		c := cell{r: 0, c: byte(j)}
		q = append(q, c)
		res[c] = true
		visited[c] = true
	}
	for i := range heights {
		c := cell{r: byte(i), c: 0}
		q = append(q, c)
		res[c] = true
		visited[c] = true
	}

	walk(m, n, q, heights, res, visited)

	return res
}

func calcAtlanticReachable(m, n int, heights [][]int) map[cell]bool {
	res := make(map[cell]bool)

	q := make([]cell, 0, m+n)
	visited := make(map[cell]bool)
	for j := range heights[0] {
		c := cell{r: byte(m - 1), c: byte(j)}
		q = append(q, c)
		res[c] = true
		visited[c] = true
	}
	for i := range heights {
		c := cell{r: byte(i), c: byte(n - 1)}
		q = append(q, c)
		res[c] = true
		visited[c] = true
	}

	walk(m, n, q, heights, res, visited)

	return res
}

func walk(m, n int, q []cell, heights [][]int, reachable, visited map[cell]bool) {
	for len(q) > 0 {
		cur := q[0]
		reachable[cur] = true
		q = q[1:]
		h := heights[cur.r][cur.c]
		if cur.c > 0 {
			l := cell{r: cur.r, c: cur.c - 1}
			if !visited[l] && heights[l.r][l.c] >= h {
				q = append(q, l)
				visited[l] = true
			}
		}
		if cur.r > 0 {
			t := cell{r: cur.r - 1, c: cur.c}
			if !visited[t] && heights[t.r][t.c] >= h {
				q = append(q, t)
				visited[t] = true
			}
		}
		if cur.c < byte(n-1) {
			r := cell{r: cur.r, c: cur.c + 1}
			if !visited[r] && heights[r.r][r.c] >= h {
				q = append(q, r)
				visited[r] = true
			}
		}
		if cur.r < byte(m-1) {
			r := cell{r: cur.r + 1, c: cur.c}
			if !visited[r] && heights[r.r][r.c] >= h {
				q = append(q, r)
				visited[r] = true
			}
		}
	}
}

type cell struct {
	r, c byte
}
