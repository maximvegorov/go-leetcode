package pacific_atlantic_water_flow

func pacificAtlantic(heights [][]int) [][]int {
	var res [][]int
	m := len(res)
	n := len(res[0])
	reachable := make([][]cell, m)
	for i := range reachable {
		reachable[i] = make([]cell, n)
	}
	for c := 0; c < n; c++ {
		for r := 0; r < m; r++ {
			if r == 0 || c == 0 {
				reachable[r][c].p = true
			} else {
				lh := heights[r][c-1]
				th := heights[r-1][c]
				h := heights[r][c]
				if reachable[r][c-1].p && lh <= h || reachable[r-1][c].p && th <= h {
					reachable[r][c].p = true
				}
			}
		}
	}

	for c := n - 1; c >= 0; c-- {
		for r := m - 1; r >= 0; r-- {
			cell := &reachable[r][c]
			if r == m-1 || c == n-1 {
				cell.a = true
			} else {
				rh := heights[r][c+1]
				bh := heights[r+1][c]
				h := heights[r][c]
				if reachable[r][c+1].a && rh <= h || reachable[r+1][c].a && bh <= h {
					cell.a = true
				}
			}
			if cell.a && cell.p {
				res = append(res, []int{r, c})
			}
		}
	}

	return res
}

type cell struct {
	p bool
	a bool
}
