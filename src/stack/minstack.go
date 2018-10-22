package stack

import (
	"math"
)

type Node struct {
	value       int
	curMinValue int
	next        *Node
	prev        *Node
}

type MinStack struct {
	ptr  *Node
	size int
}

func (this *MinStack) Push(value int) {
	/* 스택이 비어 있을 경우 */

	if this.size == 0 {
		/* value 로 노드 생성 후 스택에 추가 */
		node := Node{value, value, nil, nil}
		this.ptr = &node
	} else {
		/* 스택에 값 추가, 이 때 노드의 값은 자기 자신의 값과 현재까지 최소값 두 개의 값을 갖게 된다. */
		m := math.Min(float64(value), float64(this.GetMin()))
		node := Node{value, int(m), nil, this.ptr}
		this.ptr.next = &node
		this.ptr = &node
	}
	this.size += 1
}

func (this *MinStack) Pop() int {

	if this.size == 0 {
		return -1
	}

	this.size -= 1

	node := this.ptr
	this.ptr = node.prev

	return node.value
}

func (this *MinStack) Top() int {
	if this.size == 0 {
		return -1
	}
	return this.ptr.value
}

func (this *MinStack) GetMin() int {
	if this.size == 0 {
		return -1
	}
	return this.ptr.curMinValue
}

func (this *MinStack) Size() int {
	return this.size
}

func (this *MinStack) Empty() bool {
	return this.size == 0
}
