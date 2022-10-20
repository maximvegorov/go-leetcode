package designlinkedlist

type MyLinkedList struct {
	head         *listItem
	current      *listItem
	currentIndex int
	size         int
}

type listItem struct {
	val  int
	prev *listItem
	next *listItem
}

func (li *listItem) move(delta int) *listItem {
	p := li
	if delta >= 0 {
		for ; delta > 0; delta-- {
			p = p.next
		}
	} else {
		for ; delta < 0; delta++ {
			p = p.prev
		}
	}
	return p
}

func (li *listItem) insertBefore(val int) *listItem {
	t := &listItem{val: val, next: li, prev: li.prev}
	li.prev.next = t
	li.prev = t
	return t
}

func (li *listItem) insertAfter(val int) {
	t := &listItem{val: val, next: li.next, prev: li}
	li.next.prev = t
	li.next = t
}

func (li *listItem) delete() {
	li.next.prev = li.prev
	li.prev.next = li.next
}

func Constructor() MyLinkedList {
	head := new(listItem)
	head.next = head
	head.prev = head
	return MyLinkedList{head: head}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	this.current = this.getAtIndex(index)
	this.currentIndex = index
	return this.current.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.head.insertAfter(val)
	if this.current != nil {
		this.currentIndex++
	}
	this.size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.head.insertBefore(val)
	this.size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	this.current = this.getAtIndex(index).insertBefore(val)
	this.currentIndex = index
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	item := this.getAtIndex(index)
	if this.current != nil {
		if this.currentIndex >= index {
			if this.currentIndex == index {
				this.current = this.current.prev
			}
			this.currentIndex--
		}
	}
	item.delete()
	this.size--
	if this.size == 0 {
		this.current = nil
	}
}

func (this *MyLinkedList) getAtIndex(index int) *listItem {
	start := this.head
	delta := index + 1
	fromEnd := this.size - index
	if delta > fromEnd {
		delta = -fromEnd
	}
	if this.current != nil {
		deltaCurrent := index - this.currentIndex
		if abs(deltaCurrent) < abs(delta) {
			start = this.current
			delta = deltaCurrent
		}
	}
	return start.move(delta)
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return -x
	}
}
