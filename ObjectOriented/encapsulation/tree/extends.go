/*
go 扩展已有包：定义别名

定义别名的类型和原类型是一致的
*/

package tree

// ExtendsNode tree node的别名
type ExtendsNode = Node

// PreOrder tree node的先序遍历方法
func (node *ExtendsNode) PreOrder() {
	if node == nil {
		return
	}
	node.Print()
	node.Left.PreOrder()
	node.Right.PreOrder()
}
