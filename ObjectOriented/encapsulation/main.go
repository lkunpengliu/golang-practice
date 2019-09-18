/*
go 的封装：

go通过函数的命名来进行封装

命名的规则：
	* 名字一般使用 CamelCase
	* 首字母大写：public
	* 首字母小写：private
	* public 和 private 针对 package

包：
	* 每个目录一个包
	* main包包含可执行入口
	* 为结构定义的方法必须放在同一个包内，可以是不同的文件

扩展系统类型或其他类型的包：
	* 定义别名
	* 使用组合
*/
package main

import (
	"fmt"
	"golang-practice/ObjectOriented/encapsulation/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5, Left: nil, Right: nil}

	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Right.Left.SetValue(4)

	fmt.Println("中序遍历 -->")
	root.InOrder()
	fmt.Println()

	fmt.Println("先序遍历 -->")
	root.PreOrder()
	fmt.Println()

	fmt.Println("后序遍历 -->")
	myRoot := tree.MyTreeNode{Node: &root}
	myRoot.PostOrder()
}
