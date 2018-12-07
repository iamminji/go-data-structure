package linkedlist

import (
	"testing"
)

func TestMyLinkedList(t *testing.T) {
	obj := Constructor()

	// Node(idx:0, val: 3)
	obj.AddAtHead(3)
	t.Log(obj.String())

	// 12 => 3
	obj.AddAtHead(12)
	t.Log(obj.String())

	// 4 => 12 => 3
	obj.AddAtHead(4)
	t.Log(obj.String())

	// 4 => 12 => 3 => 11
	obj.AddAtTail(11)
	t.Log(obj.String())

	// 4 => 12 => 3 => 11 => 2
	obj.AddAtTail(2)
	t.Log(obj.String())

	// 4 => 12 => 3 => 11 => 2 => 5
	obj.AddAtTail(5)
	t.Log(obj.String())

	//	4 => 100 => 12 => 3 => 11 => 2
	obj.AddAtIndex(1, 100)
	t.Log(obj.String())

	// 99 => ...
	obj.AddAtIndex(0, 99)
	t.Log(obj.String())

	// 99 => 100 => 12 ...
	obj.DeleteAtIndex(1)
	t.Log(obj.String())

	if obj.Get(99) != -1 {
		t.Error("error")
	}

	if obj.Get(0) != 99 {
		t.Error("error")
	}
}

func TestMyLinkedList2(t *testing.T) {
	obj := Constructor()
	//
	obj.AddAtHead(1)
	t.Log(obj.Length)
	t.Log(obj.String())

	obj.AddAtTail(3)
	t.Log(obj.Length)
	t.Log(obj.String())

	obj.AddAtIndex(1, 2)
	t.Log(obj.Length)
	t.Log(obj.String())

	t.Log(obj.Get(1))

	obj.DeleteAtIndex(1)
	t.Log(obj.Get(1))
	t.Log(obj.Length)
	t.Log(obj.String())

	obj.AddAtTail(4)
	t.Log(obj.Length)
	t.Log(obj.String())
}
