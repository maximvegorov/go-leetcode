package main

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	r := rect{0, 0, len(matrix[0]), len(matrix)}
	return searchMatrixInRect(matrix, target, &r)
}

func searchMatrixInRect(matrix [][]int, target int, r *rect) bool {
	for {
		w := r.r - r.l
		h := r.b - r.t
		if w <= 0 || h <= 0 {
			return false
		} else if w == 1 && h == 1 {
			return matrix[r.t][r.l] == target
		} else if w == 1 {
			k := sort.Search(h, func(i int) bool {
				return matrix[r.t+i][r.l] >= target
			})
			return k < h && matrix[r.t+k][r.l] == target
		} else if h == 1 {
			row := matrix[r.t]
			k := sort.Search(w, func(i int) bool {
				return row[r.l+i] >= target
			})
			return k < w && row[r.l+k] == target
		}
		dx := w / 2
		dy := h / 2
		p := matrix[r.t+dy][r.l+dx]
		if target == p {
			return true
		} else if p < target {
			r1 := &rect{r.l, r.t + dy, r.r, r.b}
			r2 := &rect{r.l + dx, r.t, r.r, r.t + dy}
			if r1.weight() > r2.weight() {
				r1, r2 = r2, r1
			}
			if searchMatrixInRect(matrix, target, r1) {
				return true
			}
			r = r2
		} else {
			r1 := &rect{r.l, r.t, r.r, r.t + dy}
			r2 := &rect{r.l, r.t + dy, r.l + dx, r.b}
			if r1.weight() > r2.weight() {
				r1, r2 = r2, r1
			}
			if searchMatrixInRect(matrix, target, r1) {
				return true
			}
			r = r2
		}
	}
}

type rect struct {
	l, t, r, b int
}

func (r *rect) weight() int {
	w := r.r - r.l
	h := r.b - r.t
	if w < h {
		return w
	} else {
		return h
	}
}
