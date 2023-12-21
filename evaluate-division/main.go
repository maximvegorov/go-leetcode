package evaluatedivision

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	g := make(map[string][]*edge)
	m := make(map[equationKey]float64, len(equations))
	vars := make(map[string]bool)
	for i, eq := range equations {
		a, b := eq[0], eq[1]
		val := values[i]
		invVal := 1 / val
		g[a] = append(g[a], &edge{b, val})
		g[b] = append(g[b], &edge{a, invVal})
		m[equationKey{a, b}] = val
		m[equationKey{b, a}] = invVal
		vars[a] = true
		vars[b] = true
	}

	res := make([]float64, 0, len(queries))
	for _, query := range queries {
		a, b := query[0], query[1]
		if !vars[a] || !vars[b] {
			res = append(res, -1)
			continue
		}
		if a == b {
			res = append(res, 1)
			continue
		}
		if r, ok := m[equationKey{a, b}]; ok {
			res = append(res, r)
			continue
		}
		q := make([]queueItem, 0, 4*len(vars))
		for _, e := range g[a] {
			q = append(q, queueItem{e, e.val})
		}
		r := -1.0
		visited := map[string]bool{a: true}
		for len(q) > 0 {
			cur := q[0]
			q = q[1:]
			if cur.e.denom == b {
				r = cur.prod
				break
			}
			if visited[cur.e.denom] {
				continue
			}
			visited[cur.e.denom] = true
			for _, e := range g[cur.e.denom] {
				if visited[e.denom] {
					continue
				}
				q = append(q, queueItem{e, cur.prod * e.val})
			}
		}
		res = append(res, r)
	}
	return res
}

type edge struct {
	denom string
	val   float64
}

type equationKey struct {
	a string
	b string
}

type queueItem struct {
	e    *edge
	prod float64
}
