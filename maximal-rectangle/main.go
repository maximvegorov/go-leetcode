package maximalrectangle

func maximalRectangle(matrix [][]byte) int {
	maxSquare := 0
	var stack []PosHeight
	heights := make([]int, len(matrix[0]))
	for _, r := range matrix {
		for j, b := range r {
			if b == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
			h := heights[j]
			pos := j
			for len(stack) > 0 {
				top := &stack[len(stack)-1]
				if top.height <= h {
					break
				}
				square := (j - top.pos) * top.height
				if square > maxSquare {
					maxSquare = square
				}
				pos = top.pos
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1].height < h {
				stack = append(stack, PosHeight{pos: pos, height: h})
			}
		}
		for _, t := range stack {
			square := (len(heights) - t.pos) * t.height
			if square > maxSquare {
				maxSquare = square
			}
		}
		stack = stack[:0]
	}
	return maxSquare
}

type PosHeight struct {
	pos    int
	height int
}
