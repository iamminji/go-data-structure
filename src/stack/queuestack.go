package stack

/* Implementation Stack Using Queue */

import (
	. "queue"
)

type QueueStack struct {
	q1     *ArrayQueue
	q2     *ArrayQueue
	top    int
	q1turn bool
}

func NewQueueStack() *QueueStack {
	return new(QueueStack)
}

func (qs *QueueStack) Push(value int) {

	if qs.q1 == nil {
		qs.q1 = NewArrayQueue()
	}

	if qs.q2 == nil {
		qs.q2 = NewArrayQueue()
	}

	if qs.q1turn {
		qs.q1.Push(value)
	} else {
		qs.q2.Push(value)
	}
}

func (qs *QueueStack) Pop() int {

	if qs.q1.Empty() && qs.q2.Empty() {
		return -1
	}
	/* q1에 값이 있으면 모두 q2로 옮겨줌*/
	if qs.q1turn {
		for qs.q1.Size() > 1 {
			num := qs.q1.Pop()
			qs.q2.Push(num)
		}
		qs.q1turn = false
		return qs.q1.Pop()
	} else {
		for qs.q2.Size() > 1 {
			num := qs.q2.Pop()
			qs.q1.Push(num)
		}
		qs.q1turn = true
		return qs.q2.Pop()
	}
}

func (qs *QueueStack) Top() int {
	if qs.q1.Empty() && qs.q2.Empty() {
		return -1
	}
	/* q1에 값이 있으면 모두 q2로 옮겨줌*/
	if qs.q1turn {
		for qs.q1.Size() > 1 {
			num := qs.q1.Pop()
			qs.q2.Push(num)
		}
		t := qs.q1.Pop()
		qs.q2.Push(t)
		qs.q1turn = false
		return t
	} else {
		for qs.q2.Size() > 1 {
			num := qs.q2.Pop()
			qs.q1.Push(num)
		}
		t := qs.q2.Pop()
		qs.q1.Push(t)
		qs.q1turn = true
		return t
	}
}

func (qs *QueueStack) Empty() bool {
	return qs.q1.Empty() && qs.q2.Empty()
}
