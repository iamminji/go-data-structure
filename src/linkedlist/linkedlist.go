package linkedlist

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
}

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
		node := Node{item, nil}
		ll.head = &node
		ll.tail = &node
		ll.length = 1
	} else {
		node := Node{item, nil}
		ll.tail.next = &node
		ll.tail = ll.tail.next
		ll.length += 1
	}
}

func (ll *LinkedList) Remove(item int) {

	head := ll.head
	if head.value == item {
		ll.head = head.next
	}
	for head.next != nil {
		if head.next.value == item {
			head.next = head.next.next
			break
		}
		head = head.next
	}

}

func (ll *LinkedList) Size() int {
	return ll.length
}

func (ll *LinkedList) Pretty() {

	var str strings.Builder

	head := ll.head
	for head != nil {
		str.WriteString("Node(" + strconv.Itoa(head.value) + ")")
		str.WriteString(" -> ")
		head = head.next
	}

	str.WriteString("nil")
	fmt.Println(str.String())
}
