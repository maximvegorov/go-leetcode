package main

import "container/heap"

func leastInterval(tasks []byte, n int) int {
	if len(tasks) == 0 {
		return 0
	}
	m := make(map[byte]*taskElem)
	for _, t := range tasks {
		p := m[t]
		if p == nil {
			p = &taskElem{t: t, count: 1}
			m[t] = p
		} else {
			p.count++
		}
	}
	var unprocessed UnprocessedQueue
	for _, e := range m {
		unprocessed = append(unprocessed, e)
	}
	heap.Init(&unprocessed)
	var delayed []delayedElem
	l := 0
	for len(unprocessed) > 0 || len(delayed) > 0 {
		if len(unprocessed) == 0 {
			l = delayed[0].until
		}
		if len(delayed) > 0 && delayed[0].until <= l {
			e := delayed[0]
			delayed = delayed[1:]
			heap.Push(&unprocessed, e.t)
		}
		t := heap.Pop(&unprocessed).(*taskElem)
		t.count--
		l++
		if t.count > 0 {
			delayed = append(delayed, delayedElem{t: t, until: l + n})
		}
	}
	return l
}

type taskElem struct {
	t     byte
	count int
}

type delayedElem struct {
	t     *taskElem
	until int
}

type UnprocessedQueue []*taskElem

func (u *UnprocessedQueue) Len() int {
	return len(*u)
}

func (u *UnprocessedQueue) Less(i, j int) bool {
	return (*u)[i].count > (*u)[j].count
}

func (u *UnprocessedQueue) Swap(i, j int) {
	(*u)[i], (*u)[j] = (*u)[j], (*u)[i]
}

func (u *UnprocessedQueue) Push(x any) {
	*u = append(*u, x.(*taskElem))
}

func (u *UnprocessedQueue) Pop() any {
	l := len(*u)
	res := (*u)[l-1]
	*u = (*u)[:l-1]
	return res
}
