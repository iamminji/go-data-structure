//  online judge sites
//  https://leetcode.com/problems/design-linked-list/

package linkedlist

import (
	"fmt"
	"strings"
)

// TODO
// Add Tail O(n) -> O(1)

type MyNode struct {
	idx  int
	val  int
	next *MyNode
}

type MyLinkedList struct {
	Length int
	head   *MyNode
	tail   *MyNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	ptr := this.head
	for ptr != nil {
		if ptr.idx == index {
			return ptr.val
		}
		ptr = ptr.next
	}
	return -1
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	// 한 번도 값이 안 들어왔을 때
	if this.head == nil {
		node := &MyNode{idx: 0, val: val, next: nil}
		this.head = node
	} else {
		node := &MyNode{idx: 0, val: val, next: this.head}
		this.head = node

		// idx increment
		ptr := node.next
		for ptr != nil {
			ptr.idx += 1
			ptr = ptr.next
		}
	}
	this.Length += 1
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &MyNode{idx: this.Length, val: val, next: nil}

	ptr := this.head
	for ptr.next != nil {
		ptr = ptr.next
	}
	ptr.next = node
	this.Length += 1
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {

	if this.head == nil {
		if index == 0 {
			this.AddAtHead(val)
		}
		return
	}
	if this.Length+1 < index {
		// nothing
	} else if this.Length+1 == index {
		if this.head != nil {
			this.AddAtTail(val)
		} else {
			this.AddAtHead(val)
		}
	} else if index == 0 {
		this.AddAtHead(val)
	} else {
		node := &MyNode{idx: index, val: val}
		ptr := this.head
		var target *MyNode
		for ptr != nil {
			if ptr.idx == index-1 {
				target = ptr.next
				ptr.next = node
				node.next = target
				break
			}
			ptr = ptr.next
		}
		for target != nil {
			target.idx += 1
			target = target.next
		}
		this.Length += 1
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {

	if this.Length <= index || index < this.head.idx {
		return
	}

	this.Length -= 1

	var target *MyNode
	if index == 0 {
		target = this.head
		this.head = target.next
	} else {
		ptr := this.head
		for ptr != nil {
			if ptr.idx == index-1 {
				target = ptr.next
				if target == nil {
					// noting
				} else {
					ptr.next = target.next
				}
			}
			ptr = ptr.next
		}
	}
	for target != nil {
		target.idx -= 1
		target = target.next
	}
}

func (this *MyLinkedList) String() string {

	var str strings.Builder
	ptr := this.head
	for ptr != nil {
		str.WriteString(fmt.Sprintf("Node(idx:%d, val:%d)", ptr.idx, ptr.val))
		str.WriteString(" => ")
		ptr = ptr.next
	}
	str.WriteString("nil")
	return str.String()
}
