package main

func findBall(grid [][]int) []int {
	res := make([]int, len(grid[0]))
	for i := range res {
		x := i
		for y := 0; y < len(grid); {
			wall := grid[y][x]
			if wall == 1 {
				if x == len(res)-1 || grid[y][x+1] == -1 {
					res[i] = -1
					break
				} else {
					x++
					y++
				}
			} else if wall == -1 {
				if x == 0 || grid[y][x-1] == 1 {
					res[i] = -1
					break
				} else {
					x--
					y++
				}
			}
		}
		if res[i] != -1 {
			res[i] = x
		}
	}
	return res
}
