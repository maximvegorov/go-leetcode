package plusone

import "slices"

func plusOne(digits []int) []int {
	res := make([]int, 0, len(digits)+1)
	c := 1
	for i := len(digits) - 1; i >= 0; i-- {
		d := digits[i] + c
		if d < 10 {
			c = 0
		} else {
			d %= 10
			c = 1
		}

		res = append(res, d)
	}
	if c > 0 {
		res = append(res, c)
	}
	slices.Reverse(res)
	return res
}
