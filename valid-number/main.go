package main

func isNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	sc := newScanner([]byte(s))
	sc.matchChars('-', '+')
	if sc.matchChar('.') {
		if !sc.matchDecDigits() {
			return false
		}
		sc.skipDecDigits()
	} else if sc.matchDecDigits() {
		for sc.matchDecDigits() {
		}
		if sc.matchChar('.') {
			sc.skipDecDigits()
		}
	} else {
		return false
	}
	if sc.matchChars('e', 'E') {
		sc.matchChars('-', '+')
		if !sc.matchDecDigits() {
			return false
		}
		sc.skipDecDigits()
	}
	return sc.isEOF()
}

type scanner struct {
	token   []byte
	index   int
	curChar byte
}

func newScanner(token []byte) *scanner {
	return &scanner{
		token:   token,
		curChar: token[0],
	}
}

func (sc *scanner) matchChar(c byte) bool {
	if sc.curChar == c {
		sc.fetch()
		return true
	}
	return false
}

func (sc *scanner) matchChars(c1, c2 byte) bool {
	if sc.curChar == c1 || sc.curChar == c2 {
		sc.fetch()
		return true
	}
	return false
}

func (sc *scanner) skipDecDigits() {
	for sc.matchDecDigits() {
	}
}

func (sc *scanner) matchDecDigits() bool {
	return sc.matchRange('0', '9')
}

func (sc *scanner) matchRange(l, r byte) bool {
	if l <= sc.curChar && sc.curChar <= r {
		sc.fetch()
		return true
	}
	return false
}

func (sc *scanner) fetch() {
	sc.index++
	if sc.index < len(sc.token) {
		sc.curChar = sc.token[sc.index]
	} else {
		sc.curChar = 0
	}
}

func (sc *scanner) isEOF() bool {
	return sc.index >= len(sc.token)
}
