package queue

type PriorityQueue interface {
	Enqueue(data interface{})
	Dequeue() interface{}
	Search(data interface{}) bool
	IsEmpty() bool
	GetSize() int
}
type queue struct {
	tail *Node
	head *Node
	size int
}

func NewQueue() *queue {
	q := new(queue)
	q.head = CreateNewNode(nil)
	q.tail = q.head
	return q
}

func (q *queue) IsEmpty() bool {
	return q.head.data == nil
}

func (q *queue) Enqueue(data interface{}) {
	q.tail.data = data
	new_tail := CreateNewNode(nil)
	new_tail.next = q.tail
	q.tail.prev = new_tail
	q.tail = new_tail
	q.size++
}

func (q *queue) Dequeue() interface{} {
	if q.IsEmpty() {
		return nil
	}
	data := q.head.data
	q.head = q.head.prev
	q.head.next = nil
	q.size--
	return data
}

func (q *queue) Search(data interface{}) bool {
	for cur := q.head; cur != nil; cur = cur.prev {
		if cur.data == data {
			return true
		}
	}
	return false
}

func (q *queue) GetSize() int {
	return q.size
}
