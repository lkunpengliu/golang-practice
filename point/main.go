/*
指针

go 中的指针不能做运算
*/
package main

import "fmt"

/*
go 中的参数传递 有且只有 值传递 一种方式
*/
func passByVal(a int) {
	a++
}

/*
可以使用指针来实现 引用传递
*/
func passByRef(a *int) {
	(*a)++
}

func main() {
	a := 3
	passByVal(a)
	fmt.Println("after pass_by_val ", a)
	passByRef(&a)
	fmt.Println("after pass_by_ref", a)
}
