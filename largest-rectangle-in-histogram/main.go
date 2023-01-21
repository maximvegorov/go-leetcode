package largestrectangleinhistogram

func largestRectangleArea(heights []int) int {
	var stack []PosHeight
	maxSquare := 0
	for i, h := range heights {
		pos := i
		for len(stack) > 0 {
			top := &stack[len(stack)-1]
			if top.height <= h {
				break
			}
			square := (i - top.pos) * top.height
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
	return maxSquare
}

type PosHeight struct {
	pos    int
	height int
}
