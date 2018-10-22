package stack

import "testing"

func TestQueueStack(t *testing.T) {
	queue := NewQueueStack()

	queue.Push(3)
	queue.Push(10)
	queue.Push(1)

	if queue.Pop() != 1 {
		t.Error("error")
	}

	if queue.Pop() != 10 {
		t.Error("error")
	}

	if queue.Top() != 3 {
		t.Error("error")
	}

	queue.Push(100)

	queue.Pop()
	queue.Pop()
	queue.Push(102)

	if queue.Pop() != 102 {
		t.Error("error")
	}

	queue.Push(1)

	if queue.Top() != 1 {
		t.Error("error")
	}

	if queue.Empty() {
		t.Error("error")
	}
}
