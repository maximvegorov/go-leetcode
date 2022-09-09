package search_a_2d_matrix

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	_, found := sort.Find(m*n, func(k int) int {
		i := k / n
		j := k % n
		e := matrix[i][j]
		switch {
		case e == target:
			return 0
		case e > target:
			return -1
		default:
			return 1
		}
	})
	return found
}
