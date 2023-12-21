package integertoroman

import "strings"

func intToRoman(num int) string {
	var b strings.Builder

	num = append(&b, num, 1000, 100, 'M', 'C')
	num = append(&b, num, 500, 100, 'D', 'C')
	num = append(&b, num, 100, 10, 'C', 'X')
	num = append(&b, num, 50, 10, 'L', 'X')
	num = append(&b, num, 10, 1, 'X', 'I')
	num = append(&b, num, 5, 1, 'V', 'I')

	for i := 0; i < num; i++ {
		b.WriteByte('I')
	}

	return b.String()
}

func append(b *strings.Builder, num int, d int, step int, s1 byte, s2 byte) int {
	if num >= d {
		cnt := num / d
		for i := 0; i < cnt; i++ {
			b.WriteByte(s1)
		}
		num %= d
	}
	l := d - step
	if num >= l {
		num -= l
		b.WriteByte(s2)
		b.WriteByte(s1)
	}
	return num
}
