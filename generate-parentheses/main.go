package main

/*
func findNumberOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 0
	maxCount := 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		maxi := 0
		max0 := max
		for j := i - 1; j >= 0; j-- {
			maxj := dp[j] + 1
			if nums[j] < num && maxi < maxij {
				maxi = dp[j]
			}
			if maxi == max {
			}
		}
		maxi++
		dp[i] = maxi
		if maxi > max {
			max = maxi
			maxCount = 1
		} else if maxi == max {
			maxCount++
		}
	}
	return maxCount
}
*/

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	max := 0
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		maxi := 0
		for j := i - 1; j >= 0; j-- {
			if nums[j] <= num && maxi < dp[j] {
				maxi = dp[j]
			}
		}
		maxi++
		dp[i] = maxi
		if maxi > max {
			max = maxi
		}
	}
	return max
}

func numSquares(n int) int {
	return newSolver(n).solve(n)
}

type solver struct {
	cache []int
}

func newSolver(n int) *solver {
	cache := make([]int, n+1)
	cache2 := cache[1:]
	for i := range cache2 {
		cache2[i] = -2
	}
	return &solver{cache: cache}
}

func (s *solver) solve(n int) int {
	if res := s.cache[n]; res >= -1 {
		return res
	}
	res := n
	for i := 1; i*i <= n; i++ {
		probe := s.solve(n - i*i)
		if probe < res {
			res = probe
		}
	}
	if res < n {
		res++
	} else {
		res = -1
	}
	s.cache[n] = res
	return res
}

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	g := &generator{cache: make(map[int][]string)}
	return g.generate(n)
}

type generator struct {
	cache map[int][]string
}

func (g *generator) generate(n int) (res []string) {
	if l, ok := g.cache[n]; ok {
		return l
	}

	if n == 1 {
		res = []string{"()"}
	} else {
		for i := 1; i < n; i++ {
			left := g.generate(i)
			right := g.generate(n - i)
			set := make(map[string]struct{})
			for _, l := range left {
				for _, r := range right {
					s := l + r
					if _, found := set[s]; found {
						continue
					}
					res = append(res, s)
					set[s] = struct{}{}
				}
			}
		}
		for _, s := range g.generate(n - 1) {
			res = append(res, "("+s+")")
		}
	}

	g.cache[n] = res

	return res
}

/*
func shortestPathBinaryMatrix(grid [][]int) int {
	if grid[0][0] != 0 || grid[len(grid)-1][len(grid)-1] != 0 {
		return -1
	}
	q := []qelem{{pos: cell{r: 0, c: 0}, distance: 1}}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur.pos.r == byte(len(grid)-1) && cur.pos.c == byte(len(grid)-1) {
			return cur.distance
		}
		if cur.pos.r > 0 {
			elem := qelem{cell{r: cur.pos.r - 1, c: cur.pos.c}, cur.distance + 1}
			if grid[elem.pos.r][elem.pos.c] == 0 {
				grid[elem.pos.r][elem.pos.c] = 1
				q = append(q, elem)
			}
			if cur.pos.c > 0 {
				elem = qelem{cell{r: cur.pos.r - 1, c: cur.pos.c - 1}, cur.distance + 1}
				if grid[elem.pos.r][elem.pos.c] == 0 {
					grid[elem.pos.r][elem.pos.c] = 1
					q = append(q, elem)
				}
			}
			if cur.pos.c < byte(len(grid)-1) {
				elem = qelem{cell{r: cur.pos.r - 1, c: cur.pos.c + 1}, cur.distance + 1}
				if grid[elem.pos.r][elem.pos.c] == 0 {
					grid[elem.pos.r][elem.pos.c] = 1
					q = append(q, elem)
				}
			}
		}
		if cur.pos.c > 0 {
			elem := qelem{cell{r: cur.pos.r, c: cur.pos.c - 1}, cur.distance + 1}
			if grid[elem.pos.r][elem.pos.c] == 0 {
				grid[elem.pos.r][elem.pos.c] = 1
				q = append(q, elem)
			}
		}
		if cur.pos.c < byte(len(grid)-1) {
			elem := qelem{cell{r: cur.pos.r, c: cur.pos.c + 1}, cur.distance + 1}
			if grid[elem.pos.r][elem.pos.c] == 0 {
				grid[elem.pos.r][elem.pos.c] = 1
				q = append(q, elem)
			}
		}
		if cur.pos.r < byte(len(grid)-1) {
			elem := qelem{cell{r: cur.pos.r + 1, c: cur.pos.c}, cur.distance + 1}
			if grid[elem.pos.r][elem.pos.c] == 0 {
				grid[elem.pos.r][elem.pos.c] = 1
				q = append(q, elem)
			}
			if cur.pos.c > 0 {
				elem = qelem{cell{r: cur.pos.r + 1, c: cur.pos.c - 1}, cur.distance + 1}
				if grid[elem.pos.r][elem.pos.c] == 0 {
					grid[elem.pos.r][elem.pos.c] = 1
					q = append(q, elem)
				}
			}
			if cur.pos.c < byte(len(grid)-1) {
				elem = qelem{cell{r: cur.pos.r + 1, c: cur.pos.c + 1}, cur.distance + 1}
				if grid[elem.pos.r][elem.pos.c] == 0 {
					grid[elem.pos.r][elem.pos.c] = 1
					q = append(q, elem)
				}
			}
		}
	}
	return -1
}

type cell struct {
	r, c byte
}

type qelem struct {
	pos      cell
	distance int
}

/*
type trieElem struct {
	children map[byte]*trieElem
}

func findWords(board [][]byte, words []string) []string {
	trie := make(map[byte][]cell)
	for i, column := range board {
		for j, c := range column {
			trie[c] = append(trie[c], cell{r: byte(i), c: byte(j)})
		}
	}
	m := len(board)
	n := len(board[0])
	res := []string{}
	for _, word := range words {
		chars := []byte(word)
		startCells, found := trie[chars[0]]
		if !found {
			continue
		}
	outer:
		for _, startCell := range startCells {
			q := []*qelem{{next: startCell}}
			for len(q) > 0 {
				cur := q[0]
				q = q[1:]
				pos := cur.next
				if board[pos.r][pos.c] != chars[cur.matchedCount] {
					continue
				}
				used := make(map[cell]struct{})
				for e := cur.matched; e != nil; e = e.next {
					used[e.pos] = struct{}{}
				}
				used[pos] = struct{}{}
				matchedCount := cur.matchedCount + 1
				if int(matchedCount) == len(chars) {
					res = append(res, word)
					break outer
				}
				matched := &matchedSeq{pos, cur.matched}
				if pos.r > 0 {
					u := cell{pos.r - 1, pos.c}
					if _, found := used[u]; !found {
						qe := &qelem{next: u, matched: matched, matchedCount: matchedCount}
						q = append(q, qe)
					}
				}
				if pos.c > 0 {
					l := cell{pos.r, pos.c - 1}
					if _, found := used[l]; !found {
						qe := &qelem{next: l, matched: matched, matchedCount: matchedCount}
						q = append(q, qe)
					}
				}
				if int(pos.r) < m-1 {
					b := cell{pos.r + 1, pos.c}
					if _, found := used[b]; !found {
						qe := &qelem{next: b, matched: matched, matchedCount: matchedCount}
						q = append(q, qe)
					}
				}
				if int(pos.c) < n-1 {
					r := cell{pos.r, pos.c + 1}
					if _, found := used[r]; !found {
						qe := &qelem{next: r, matched: matched, matchedCount: matchedCount}
						q = append(q, qe)
					}
				}
			}
		}
	}
	return res
}

type qelem struct {
	next         cell
	matched      *matchedSeq
	matchedCount byte
}

type matchedSeq struct {
	pos  cell
	next *matchedSeq
}

type cell struct {
	r, c byte
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	if numCourses == 0 {
		return true
	}
	deps := make([][]int, numCourses)
	for _, p := range prerequisites {
		course := p[0]
		deps[course] = append(deps[course], p[1])
	}
	ctx := canFinishContext{deps: deps, finishable: make(map[int]int)}
	for i := range deps {
		if !canFinishCourse(&ctx, i) {
			return false
		}
	}
	return true
}

const (
	IN_PROGRESS = iota
	FINISHABLE
)

type canFinishContext struct {
	deps       [][]int
	finishable map[int]int
}

func canFinishCourse(ctx *canFinishContext, course int) bool {
	if len(ctx.deps[course]) > 0 {
		for _, dep := range ctx.deps[course] {
			state, ok := ctx.finishable[dep]
			if ok {
				if state == IN_PROGRESS {
					return false
				}
			} else {
				ctx.finishable[dep] = IN_PROGRESS
				if !canFinishCourse(ctx, dep) {
					return false
				}
			}
		}
	}
	ctx.finishable[course] = FINISHABLE
	return true
}

func intervalIntersection(f [][]int, s [][]int) [][]int {
	if len(f) == 0 || len(s) == 0 {
		return [][]int{}
	}
	var res [][]int
	i, j := 0, 0
	for i < len(f) && j < len(s) {
		lf, rf := f[i][0], f[i][1]
		ls, rs := s[j][0], s[j][1]
		if lf <= ls {
			if rf >= ls {
				res = append(res, []int{maxOf(lf, ls), minOf(rf, rs)})
			}
			if rf <= rs {
				i++
			} else {
				j++
			}
		} else {
			if rs >= lf {
				res = append(res, []int{maxOf(lf, ls), minOf(rf, rs)})
			}
			if rs <= rf {
				j++
			} else {
				i++
			}
		}
	}
	return res
}

func backspaceCompare(s string, t string) bool {
	b1 := processBackspaces(s)
	b2 := processBackspaces(t)
	return bytes.Equal(b1, b2)
}

func processBackspaces(s string) []byte {
	var b []byte
	for _, c := range []byte(s) {
		if c == '#' {
			if len(b) > 0 {
				b = b[:len(b)-1]
			}
		} else {
			b = append(b, c)
		}
	}
	return b
}

func trailingZeroes(n int) int {
	var count2, count5 int
	for i := 2; i <= n; i++ {
		count2 += countOf(i, 2)
		count5 += countOf(i, 5)
	}
	return minOf(count2, count5)
}

func countOf(n, factor int) (count int) {
	for n%factor == 0 {
		count++
		n /= factor
	}
	return
}

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for r-l >= 2 {
		m := (l + r) / 2
		if nums[m] < nums[m+1] {
			l = m
		} else if nums[m] < nums[m-1] {
			r = m
		} else {
			return m
		}
	}
	switch r - l {
	case 0:
		return nums[l]
	case 1:
		if nums[l] < nums[r] {
			return nums[r]
		} else {
			return nums[l]
		}
	default:
		panic("Error")
	}
}

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	subList := indexOfLesser(nums)
	if subList == nil {
		return nums[0]
	}
	return extractOfMin(subList)
}

func indexOfLesser(nums []int) []int {
	l, r := 0, len(nums)
	for l < r-1 {
		m := (l + r) / 2
		if nums[m] > nums[l] {
			l = m
		} else if nums[m] < nums[l] {
			return nums[l : m+1]
		}
	}
	return nil
}

func extractOfMin(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	l, r := 0, len(nums)-1
	for l != r-1 {
		m := (l + r) / 2
		if nums[m] > nums[l] {
			l = m
		} else if nums[m] < nums[r] {
			r = m
		}
	}
	return nums[r]
}

func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}
	freqs := make(map[byte]int)
	for _, t := range tasks {
		freqs[t] += 1
	}
	maxFreq := 0
	for _, freq := range freqs {
		if freq > maxFreq {
			maxFreq = freq
		}
	}
	maxFreqCount := 0
	for _, freq := range freqs {
		if freq == maxFreq {
			maxFreqCount++
		}
	}
	return maxFreqCount

}

func longestPalindrome(words []string) int {
	counts := map[[2]byte]int{}
	for _, w := range words {
		bytes := []byte(w)
		k := [...]byte{bytes[0], bytes[1]}
		counts[k] += 1
	}
	maxLength := 0
	hasOddDouble := false
	for _, w := range words {
		bytes := []byte(w)
		k := [...]byte{bytes[0], bytes[1]}
		if k[0] == k[1] {
			used := counts[k]
			maxLength += (used & ^1) * len(k)
			delete(counts, k)
			if !hasOddDouble && (used&1) != 0 {
				maxLength += len(k)
				hasOddDouble = true
			}
		} else {
			invK := [...]byte{bytes[1], bytes[0]}
			used := minOf(counts[k], counts[invK])
			maxLength += 2 * used * len(k)
			counts[k] -= used
			counts[invK] -= used
		}
	}
	return maxLength
}

func makeList(nums ...int) *ListNode {
	head := ListNode{}
	p := &head
	for _, v := range nums {
		q := &ListNode{Val: v}
		p.Next = q
		p = q
	}
	return head.Next
}

func printList(head *ListNode) {
	for p := head; p != nil; p = p.Next {
		fmt.Println(p.Val)
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	oddHead, p := head, head
	evenHead, q := head.Next, head.Next

	for {
		if q.Next == nil {
			break
		}
		p.Next = q.Next
		p = p.Next

		if p.Next == nil {
			q.Next = nil
			break
		}

		q.Next = p.Next
		q = q.Next
	}

	p.Next = evenHead

	return oddHead
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l, r := splitList(head)

	l = sortList(l)
	r = sortList(r)

	sortedHead := ListNode{}
	p := &sortedHead
	for l != nil && r != nil {
		if l.Val < r.Val {
			p.Next = l
			l = l.Next
		} else {
			p.Next = r
			r = r.Next
		}
		p = p.Next
	}
	if l != nil {
		p.Next = l
	} else if r != nil {
		p.Next = r
	}
	return sortedHead.Next
}

func splitList(head *ListNode) (left *ListNode, right *ListNode) {
	p, q := &ListNode{Next: head}, head.Next
	for q != nil {
		p = p.Next
		q = q.Next
		if q == nil {
			break
		}
		q = q.Next
	}
	left = head
	right = p.Next
	p.Next = nil
	return
}

func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	startIndex := 0
	window := []byte{s[0]}
	freqs := map[byte]int{s[0]: 1}
	max := 1
	for i := 1; i < len(s); i++ {
		ci := s[i]
		freqs[ci]++
		window = append(window, ci)
		if ci == window[0] || len(window)-freqs[window[0]] <= k {
			continue
		}
		max = maxOf(max, len(window)-1)
		for len(window)-freqs[window[0]] > k {
			for j, cj := range window {
				if cj != window[0] {
					startIndex += j
					freqs[window[0]] -= j
					window = window[j:]
					break
				}
			}
		}
	}
	replacedChar := len(window) - freqs[window[0]]
	max = maxOf(max, minOf(startIndex, k-replacedChar)+len(window))
	return max
}

func minOf(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func maxOf(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
*/
