package dinnerplatestacks

type DinnerPlates struct {
	stacks []*Stack
	maxCap int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{stacks: []*Stack{newStack(capacity)}, maxCap: capacity}
}

func (this *DinnerPlates) Push(val int) {
	curStack := this.stacks[len(this.stacks)-1]
	if curStack.isFull() {
		curStack = newStack(this.maxCap)
		this.stacks = append(this.stacks, curStack)
		
	}
	curStack.push(val)
}

func (this *DinnerPlates) Pop() int {
	curStack := this.stacks[len(this.stacks)-1]
	for curStack.isEmpty() {
		if len(this.stacks) == 1 {
			return -1
		}
		this.stacks = this.stacks[:len(this.stacks)-1]
		curStack = this.stacks[len(this.stacks)-1]
	}
	return curStack.pop()
}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index < 0 || index >= len(this.stacks) {
		return -1
	}
	return this.stacks[index].pop()
}

type Stack struct {
	items []int
}

func newStack(maxCap int) *Stack {
	return &Stack{items: make([]int, 0, maxCap)}
}

func (s *Stack) push(val int) {
	s.items = append(s.items, val)
}

func (s *Stack) pop() int {
	if len(s.items) == 0 {
		return -1
	}
	res := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return res
}

func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) isFull() bool {
	return len(s.items) == cap(s.items)
}
