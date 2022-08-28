package main

func isHappy(n int) bool {
	generated := make(map[int]bool, 729)
	for !generated[n] {
		generated[n] = true
		n = sumOfDigitSquares(n)
	}
	return n == 1
}

func sumOfDigitSquares(n int) int {
	res := 0
	for n > 0 && res >= 0 {
		d := n % 10
		n /= 10
		res += d * d
	}
	return res
}
