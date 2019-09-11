/*
数组

* 数组是值类型
	[10]int 和 [20]int 是不同的类型
	调用 func f(arr [10]int) 会 拷贝 数组
*/
package main

import "fmt"

func main() {
	// 数组定义方法一： 直接定义，可以没有初始值
	var arr1 [5]int
	// 数组定义方法二：用 := 定义，必须有初始值
	arr2 := [3]int{1, 3, 5}
	// 数组定义方法三：没有数组长度，必须有 ...
	arr3 := [...]int{2, 4, 6, 8}
	// 二维数组定义方法
	var grid [3][4]bool

	fmt.Println(
		arr1,
		arr2,
		arr3,
	)
	fmt.Println(grid)

	// 数组遍历方法一： 使用for循环
	for i := 0; i < len(arr3); i++ {
		fmt.Println(arr3[i])
	}
	// 数组遍历方法二：使用range关键字
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
}
