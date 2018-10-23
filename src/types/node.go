package types

import "errors"

type Node struct {
	Item interface{}
	Prev *Node
	Next *Node
}

var errorNewNode = errors.New("failed to create node")

func NewNode(item interface{}) (*Node, error) {
	node := &Node{item, nil, nil}
	if node == nil {
		return nil, errorNewNode
	}
	return node, nil
}
