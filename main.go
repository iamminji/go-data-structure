package main

import (
	"fmt"
	. "types"
)

func main() {
	fmt.Println("Test")
	node, _ := NewNode(2)
	fmt.Println(node)

	node2, _ := NewNode("Hello")
	fmt.Println(node2)
}
