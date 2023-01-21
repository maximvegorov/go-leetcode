package designfrontmiddlebackqueue

type FrontMiddleBackQueue struct {
	left  *QueueItem
	right *QueueItem
	size  int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{left: newQueue(), right: newQueue()}
}

func (q *FrontMiddleBackQueue) PushFront(val int) {
	q.left.insertAfter(&QueueItem{val: val})
	q.size++
	if q.size%2 != 0 {
		last := q.left.prev.remove()
		q.right.insertAfter(last)
	}
}

func (q *FrontMiddleBackQueue) PushMiddle(val int) {
	q.right.insertAfter(&QueueItem{val: val})
	q.size++
	if q.size%2 != 0 {
		last := q.left.prev.remove()
		q.right.insertAfter(last)
	}
}

func (q *FrontMiddleBackQueue) PushBack(val int) {

}

func (q *FrontMiddleBackQueue) PopFront() int {

}

func (this *FrontMiddleBackQueue) PopMiddle() int {

}

func (this *FrontMiddleBackQueue) PopBack() int {

}

type QueueItem struct {
	val  int
	prev *QueueItem
	next *QueueItem
}

func newQueue() *QueueItem {
	q := &QueueItem{}
	q.next = q
	q.prev = q
	return q
}
func (qi *QueueItem) insertBefore(newItem *QueueItem) {
	newItem.next = qi
	newItem.prev = qi.prev
	qi.prev.next = newItem
	qi.prev = newItem
}

func (qi *QueueItem) insertAfter(newItem *QueueItem) {
	newItem.prev = qi
	newItem.next = qi.prev
	qi.next.prev = newItem
	qi.next = newItem
}

func (qi *QueueItem) remove() *QueueItem {
	qi.next.prev = qi.prev
	qi.prev.next = qi.next
}
