package increasingtripletsubsequence

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	p0, i := findPair(&pair{}, nums, 0)
	for i < len(nums) {
		if nums[i] > p0.b {
			return true
		}
		var p1 pair
		p1, i = findPair(&p0, nums, i)
		var found bool
		p0, found = mergePair(&p0, &p1)
		if found {
			return true
		}
	}
	return false
}

func mergePair(p0 *pair, p1 *pair) (pair, bool) {
	if !p1.isValid() {
		return *p0, false
	}
	if p0.b < p0.a {
		return *p0, true
	} else if p0.a < p1.a {
		return pair{a: p0.a, b: p1.a}, true
	} else if p0.b < p1.b {
		return *p0, true
	} else {
		return *p1, false
	}
}

func findPair(p *pair, nums []int, start int) (pair, int) {
	i := start
	for i < len(nums)-1 && nums[i] >= nums[i+1] {
		num := nums[i]
		if p.a < num {
			p.b = num
		}
		i++
	}
	if i >= len(nums)-1 {
		t := nums[len(nums)-1]
		return pair{a: t, b: t}, len(nums)
	}
	return pair{a: nums[i], b: nums[i+1]}, i + 2
}

type pair struct {
	a, b int
}

func (p *pair) isValid() bool {
	return p.a < p.b
}
