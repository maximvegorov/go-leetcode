package checkifnumberisasumofpowersofthree

func checkPowersOfThree(n int) bool {
	for n != 0 {
		d := n % 3
		if d > 1 {
			return false
		}
		n = n / 3
	}
	return true
}
