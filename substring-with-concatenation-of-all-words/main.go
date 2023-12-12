package substringwithconcatenationofallwords

func findSubstring(s string, words []string) []int {
	dict := newWordSet()
	for _, w := range words {
		dict.add(w)
	}

	res := make([]int, 0, 1)
	curset := newWordSet()
	for i, n := 0, len(words[0]); i < n; i++ {
		curset.clear()
		for j := i; j+n <= len(s); j += n {
			w := s[j : j+n]
			targetCnt := dict.count(w)
			if targetCnt == 0 {
				curset.clear()
				continue
			}
			if curCnt := curset.count(w); curCnt == targetCnt {
				for k := j - curset.size*n; k < j; k += n {
					t := s[k : k+n]
					curset.remove(t)
					if t == w {
						break
					}
				}
			}
			curset.add(w)
			if curset.size == dict.size {
				res = append(res, j-curset.size*n+n)
			}
		}
	}
	return res
}

type wordSet struct {
	freqs map[string]int
	size  int
}

func newWordSet() *wordSet {
	return &wordSet{freqs: make(map[string]int)}
}

func (ws *wordSet) count(w string) int {
	return ws.freqs[w]
}

func (ws *wordSet) add(w string) {
	ws.freqs[w] += 1
	ws.size += 1
}

func (ws *wordSet) remove(w string) {
	ws.freqs[w] -= 1
	ws.size -= 1
}

func (ws *wordSet) clear() {
	clear(ws.freqs)
	ws.size = 0
}
