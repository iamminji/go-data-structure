package linkedlist

import (
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	ll := NewLinkedList()
	ll.Add(3)
	ll.Add(1)
	ll.Add(2)

	ll.Pretty()

}

func TestLinkedList_Remove(t *testing.T) {
	ll := NewLinkedList()
	ll.Add(3)
	ll.Add(2)
	ll.Add(10)

	ll.Remove(2)
	ll.Pretty()

	ll.Add(9)
	ll.Add(8)
	ll.Add(7)
	ll.Add(6)

	ll.Pretty()
	ll.Remove(10)
	ll.Remove(3)
	ll.Remove(6)

	ll.Pretty()
}
