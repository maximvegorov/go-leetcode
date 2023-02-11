package finduniquebinarystring

func findDifferentBinaryString(bwords []string) string {
	set := toSet(bwords)
	n := len(bwords[0])
	num := 0
	for set[num] {
		num++
	}
	return toBword(num, n)
}

func toSet(bwords []string) map[int]bool {
	set := make(map[int]bool, len(bwords))
	for _, bword := range bwords {
		num := 0
		for i, b := range bword {
			if b == '1' {
				num |= 1 << i
			}
		}
		set[num] = true
	}
	return set
}

func toBword(num int, n int) string {
	bword := make([]byte, n)
	for i := 0; i < n; i++ {
		if (num & (1 << i)) != 0 {
			bword[i] = '1'
		} else {
			bword[i] = '0'
		}
	}
	return string(bword)
}
