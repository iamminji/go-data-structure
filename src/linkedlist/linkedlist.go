package linkedlist

import (
	"fmt"
	"strconv"
	"strings"
	. "types"
)

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewLinkedList() *LinkedList {
	return new(LinkedList)
}

func (ll *LinkedList) Add(item int) {

	if ll.head == nil {
		node, _ := NewNode(item)
		ll.head = node
		ll.tail = node
		ll.length = 1
	} else {
		node, _ := NewNode(item)
		ll.tail.Next = node
		ll.tail = ll.tail.Next
		ll.length += 1
	}
}

func (ll *LinkedList) Remove(item int) {

	head := ll.head
	if head.Item == item {
		ll.head = head.Next
	}
	for head.Next != nil {
		if head.Next.Item == item {
			head.Next = head.Next.Next
			break
		}
		head = head.Next
	}

}

func (ll *LinkedList) Size() int {
	return ll.length
}

func (ll *LinkedList) Pretty() {

	var str strings.Builder

	head := ll.head
	for head != nil {
		switch head.Item.(type) {
		case string:
			str.WriteString("Node(" + head.Item.(string) + ") -> ")
		case int:
			str.WriteString("Node(" + strconv.Itoa(head.Item.(int)) + ") -> ")
		default:
			str.WriteString("Unknown Type")
		}

		head = head.Next
	}

	str.WriteString("nil")
	fmt.Println(str.String())
}
