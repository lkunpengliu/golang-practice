package tree

import "fmt"

type Node struct {
	Value       int
	Left, Right *Node
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil node. Ignored.")
		return
	}
	node.Value = value
}

func (node Node) Print() {
	fmt.Printf("%d ", node.Value)
}

func CreateNode(value int) *Node {
	return &Node{value, nil, nil}
}
