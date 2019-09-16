/*
go 分支条件语句
1. if语句
	* 不需要用括号来包住条件
	* 条件可以是多条语句
	* 条件中定义的变量的作用域只能在if块中使用
2. switch语句
	* 不需要添加 break
	* switch 后可以没有条件语句
*/
package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		// panic 可以中断代码的执行
		panic(fmt.Sprintf("Wrong score: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s", contents)
	}

	fmt.Println(
		grade(0),
		grade(59),
		grade(65),
		grade(82),
		grade(99),
		grade(100),
	)
}
