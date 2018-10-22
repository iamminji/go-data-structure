package queue

import "testing"

func TestArrayQueue_Push(t *testing.T) {
	queue := NewArrayQueue()

	queue.Push(3)
	queue.Push(2)
}

func TestArrayQueue_Pop(t *testing.T) {
	queue := NewArrayQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	if queue.Pop() != 1 {
		t.Error("error")
	}
	queue.Push(100)

	if queue.Pop() != 2 {
		t.Error("error")
	}
}

func TestArrayQueue_Top(t *testing.T) {
	queue := NewArrayQueue()
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)

	if queue.Top() != 1 {
		t.Error("error")
	}

	queue.Pop()

	if queue.Top() != 2 {
		t.Error("error")
	}
}

func TestArrayQueue_Empty(t *testing.T) {
	queue := NewArrayQueue()
	if !queue.Empty() {
		t.Error("error")
	}
}
