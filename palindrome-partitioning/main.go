package main

func partition(str string) [][]string {
	s := NewPartitionSearcher([]byte(str))
	chains := s.search(0)
	var res [][]string
	for _, chain := range chains {
		res = append(res, chain.toSlice())
	}
	return res
}

type partitionSearcher struct {
	s         []byte
	palidroms [][]int
	cache     [][]*partitionChainElem
}

func NewPartitionSearcher(s []byte) *partitionSearcher {
	isPalidrom := make([][]bool, len(s))
	for i := range isPalidrom {
		isPalidrom[i] = make([]bool, len(s))
	}
	for l := 0; l < len(s); l++ {
		for i := 0; i < len(s)-l; i++ {
			j := i + l
			if s[i] != s[j] {
				continue
			}
			isPalidrom[i][j] = i == j || j == i+1 || isPalidrom[i+1][j-1]
		}
	}
	palidroms := make([][]int, len(s))
	for i := range isPalidrom {
		r := isPalidrom[i]
		for j := i; j < len(r); j++ {
			if r[j] {
				palidroms[i] = append(palidroms[i], j+1)
			}
		}
	}
	return &partitionSearcher{
		s:         []byte(s),
		palidroms: palidroms,
		cache:     make([][]*partitionChainElem, len(s)),
	}
}

func (s *partitionSearcher) search(start int) []*partitionChainElem {
	if start >= len(s.s) {
		return nil
	}
	if e := s.cache[start]; e != nil {
		return e
	}
	var res []*partitionChainElem
	for _, j := range s.palidroms[start] {
		tails := s.search(j)
		subStr := s.s[start:j]
		if tails == nil {
			res = append(res, &partitionChainElem{v: subStr})
			continue
		}
		for _, tail := range tails {
			res = append(res, &partitionChainElem{v: subStr, prev: tail})
		}
	}
	s.cache[start] = res
	return res
}

type partitionChainElem struct {
	v    []byte
	prev *partitionChainElem
}

func (h *partitionChainElem) toSlice() []string {
	var res []string
	for e := h; e != nil; e = e.prev {
		res = append(res, string(e.v))
	}
	return res
}
