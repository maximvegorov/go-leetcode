package addbinary

import "slices"

func addBinary(a string, b string) string {
	n := min(len(a), len(b))
	var acc accumulator
	for i := 1; i <= n; i++ {
		d1 := atoi(a[len(a)-i])
		d2 := atoi(b[len(b)-i])
		acc.append(d1 + d2)
	}
	for i := n + 1; i <= len(a); i++ {
		acc.append(atoi(a[len(a)-i]))
	}
	for i := n + 1; i <= len(b); i++ {
		acc.append(atoi(b[len(b)-i]))
	}
	return acc.Finish()
}

type accumulator struct {
	buf   []byte
	carry byte
}

func (a *accumulator) append(d byte) {
	d += a.carry
	if d > 1 {
		d %= 2
		a.carry = 1
	} else {
		a.carry = 0
	}
	if d == 0 {
		a.buf = append(a.buf, '0')
	} else {
		a.buf = append(a.buf, '1')
	}
}

func (a *accumulator) Finish() string {
	if a.carry > 0 {
		a.buf = append(a.buf, '1')
	}
	slices.Reverse(a.buf)
	return string(a.buf)
}

func atoi(d byte) (res byte) {
	if d == '1' {
		res = 1
	}
	return
}
