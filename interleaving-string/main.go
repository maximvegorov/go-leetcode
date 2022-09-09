package interleaving_string

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	s := newSolver(s3, s1, s2)
	return s.solve(0, 0)
}

type SolveStatus byte

const (
	None SolveStatus = iota
	True
	False
)

type solver struct {
	cache [][]SolveStatus
	s     string
	s1    string
	s2    string
}

func newSolver(s string, s1 string, s2 string) solver {
	cache := make([][]SolveStatus, len(s1)+1)
	for i := range cache {
		cache[i] = make([]SolveStatus, len(s2)+1)
	}
	return solver{
		cache: cache,
		s:     s,
		s1:    s1,
		s2:    s2,
	}
}

func (s *solver) solve(k1, k2 int) bool {
	if res := s.cache[k1][k2]; res != None {
		return res == True
	}
	l := k1 + k2
	if k1 == len(s.s1) {
		var res SolveStatus
		if s.s[l:] == s.s2[k2:] {
			res = True
		} else {
			res = False
		}
		s.cache[k1][k2] = res
		return res == True
	}
	if k2 == len(s.s2) {
		var res SolveStatus
		if s.s[l:] == s.s1[k1:] {
			res = True
		} else {
			res = False
		}
		s.cache[k1][k2] = res
		return res == True
	}
	for i := 0; i < len(s.s1)-k1 && s.s1[k1+i] == s.s[l+i]; i++ {
		if s.solve(k1+i+1, k2) {
			s.cache[k1][k2] = True
			return true
		}
	}
	for i := 0; i < len(s.s2)-k2 && s.s2[k2+i] == s.s[l+i]; i++ {
		if s.solve(k1, k2+i+1) {
			s.cache[k1][k2] = True
			return true
		}
	}
	s.cache[k1][k2] = False
	return false
}
