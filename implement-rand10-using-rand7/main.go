package implementrand10usingrand7

func rand7() int {
	return 0
}

func rand10() int {
	for {
		d0 := rand7() - 1
		d1 := rand7() - 1
		d2 := rand7() - 1
		d3 := rand7() - 1
		n := 7*(7*(7*d3+d2)+d1) + d0
		if n < 2400 {
			return n/240 + 1
		}
	}
}
