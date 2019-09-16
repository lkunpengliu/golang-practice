package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

/*
结构体定义方法：func (方法接收者) 方法名
	* 方法接受者是值传递，只有使用指针才能改变结构内容
	* nil 指针也可以调用方法
*/
// 为结构定义方法 显示定义和命名方法接收者
func (node treeNode) print() {
	fmt.Print(node.value)
}

// 使用指针作为方法接受者
func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func createNode(value int) *treeNode {
	return &treeNode{value, nil, nil}
}

func main() {
	var root treeNode
	// 使用结构体直接构造 treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	// 使用new关键字构造 treeNode
	root.right.left = new(treeNode)
	// 使用工厂函数构造 treeNode
	root.left.right = createNode(2)

	fmt.Println("root ->", root)

	root.print()
	root.right.left.setValue(4)
	root.right.left.print()
	fmt.Println("root 的中序遍历 -->")
	root.traverse()
}
