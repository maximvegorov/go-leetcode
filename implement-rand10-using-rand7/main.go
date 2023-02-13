package implementrand10usingrand7

func rand7() int {
	return 0
}

func rand10() int {
	for {
		h := rand7() - 1
		l := rand7() - 1
		n := h*7 + l
		if n <= 9 {
			return n + 1
		}
	}
}
