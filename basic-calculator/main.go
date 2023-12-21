package basiccalculator

import (
	"bytes"
	"io"
	"strconv"
	"strings"
)

func calculate(s string) int {
	scan := newScanner(s)
	return calcExpr(&scan)
}

func calcExpr(s *scanner) int {
	var res int
	var op byte = '+'
	for {
		neg := s.scanMinus()
		var rval int
		if s.scanLParentheses() {
			rval = calcExpr(s)
		} else {
			rval = s.scanNumber()
		}
		if neg {
			rval = -rval
		}
		switch op {
		case '+':
			res += rval
		case '-':
			res -= rval
		}
		op = s.scanOp()
		if op == ')' {
			return res
		}
	}
}

type scanner struct {
	r io.ByteScanner
	b *bytes.Buffer
}

func newScanner(s string) scanner {
	return scanner{strings.NewReader(s), bytes.NewBuffer(make([]byte, 16))}
}

func (s *scanner) skipSpaces() (byte, error) {
	c, err := s.r.ReadByte()
	for err != io.EOF && c == ' ' {
		c, err = s.r.ReadByte()
	}
	return c, err
}

func (s *scanner) scanMinus() bool {
	c, _ := s.skipSpaces()

	if c == '-' {
		return true
	}

	s.r.UnreadByte()

	return false
}

func (s *scanner) scanLParentheses() bool {
	c, _ := s.skipSpaces()

	if c == '(' {
		return true
	}

	s.r.UnreadByte()

	return false
}

func (s *scanner) scanNumber() int {
	c, _ := s.skipSpaces()

	s.b.Reset()
	s.b.WriteByte(c)
	c, err := s.r.ReadByte()
	for '0' <= c && c <= '9' {
		s.b.WriteByte(c)
		c, err = s.r.ReadByte()
	}
	if err != io.EOF {
		s.r.UnreadByte()
	}

	num, _ := strconv.Atoi(s.b.String())

	return num
}

func (s *scanner) scanOp() byte {
	c, err := s.skipSpaces()
	if err == io.EOF {
		return ')'
	}
	return c
}
