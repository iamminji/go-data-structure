package queue

type ArrayQueue struct {
	array []int
}

func NewArrayQueue() *ArrayQueue {
	nq := new(ArrayQueue)
	return nq
}

func (q *ArrayQueue) Push(value int) {
	q.array = append(q.array, value)
}

func (q *ArrayQueue) Pop() int {

	num := q.array[0]
	q.array = q.array[1:]
	return num
}

func (q *ArrayQueue) Top() int {
	if len(q.array) == 0 {
		return -1
	}

	return q.array[0]
}

func (q *ArrayQueue) Size() int {
	return len(q.array)
}

func (q *ArrayQueue) Empty() bool {
	return len(q.array) == 0
}
