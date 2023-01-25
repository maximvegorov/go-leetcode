package stackofstacks

import "fmt"

type StackOfStacks[T any] struct {
	stack   [][]T
	maxSize int
}

func NewStackOfStacks[T any](maxSize int) *StackOfStacks[T] {
	if maxSize <= 0 {
		panic(fmt.Errorf("invalid max size %v", maxSize))
	}
	stack := [][]T{make([]T, 0, maxSize)}
	return &StackOfStacks[T]{
		stack:   stack,
		maxSize: maxSize,
	}
}

func (ss *StackOfStacks[T]) Push(elem T) {
	current := &ss.stack[len(ss.stack)-1]
	if len(*current) >= ss.maxSize {
		ss.stack = append(ss.stack, make([]T, 0, ss.maxSize))
		current = &ss.stack[len(ss.stack)-1]
	}
	*current = append(*current, elem)
}

func (ss *StackOfStacks[T]) Pop() (T, error) {
	var zero T
	current := &ss.stack[len(ss.stack)-1]
	if len(*current) == 0 {
		return zero, fmt.Errorf("stack is empty")
	}
	p := &((*current)[len(*current)-1])
	res := *p
	*p = zero
	*current = (*current)[:len(*current)-1]
	if len(*current) == 0 && len(ss.stack) > 1 {
		ss.stack = ss.stack[:len(ss.stack)-1]
	}
	return res, nil
}
