package tree

// InOrder tree node的中序遍历方法
func (node *Node) InOrder() {
	if node == nil {
		return
	}
	node.Left.InOrder()
	node.Print()
	node.Right.InOrder()
}
