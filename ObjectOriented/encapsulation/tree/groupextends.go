/*
go 扩展已有包：组合方式

组合方式中 MyTreeNode 和原类型 Node 是两个完全不同的类型
*/

package tree

// MyTreeNode 新的类型，Node继承于 tree node
type MyTreeNode struct {
	Node *Node
}

// PostOrder tree node的后序遍历方法
func (myNode *MyTreeNode) PostOrder() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	/*
		错误写法：会出现 cannot call pointer method on myTreeNode literal 错误提示
		错误原因：指针不能作为接收者，需要用定义变量来接送地址
		myTreeNode{myNode.node.Left}.postOrder()
		myTreeNode{myNode.node.Right}.postOrder()
		myNode.node.Print()
	*/
	left := MyTreeNode{myNode.Node.Left}
	left.PostOrder()
	right := MyTreeNode{myNode.Node.Right}
	right.PostOrder()
	myNode.Node.Print()
}
