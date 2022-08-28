package main

func minDistance(word1 string, word2 string) int {
	s1 := []byte(word1)
	s2 := []byte(word2)
	if len(s1) == 0 {
		return len(s2)
	} else if len(s2) == 0 {
		return len(s1)
	}
	s := newSolver(s1, s2)
	return s.solve(len(s1)-1, len(s2)-1)
}

type solver struct {
	cache [][]int
	s1    []byte
	s2    []byte
}

func newSolver(s1 []byte, s2 []byte) *solver {
	cache := make([][]int, len(s1))
	for i := range cache {
		cache[i] = make([]int, len(s2))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return &solver{cache: cache, s1: s1, s2: s2}
}

func (s *solver) solve(i, j int) int {
	if i < 0 {
		return j + 1
	} else if j < 0 {
		return i + 1
	}
	if r := s.cache[i][j]; r >= 0 {
		return r
	}
	r := s.solve(i-1, j) + 1
	r = minOf(r, s.solve(i, j-1)+1)
	d := s.solve(i-1, j-1)
	if s.s1[i] != s.s2[j] {
		d++
	}
	r = minOf(r, d)
	s.cache[i][j] = r
	return r
}

func minOf(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
