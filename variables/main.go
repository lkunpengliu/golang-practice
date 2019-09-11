/*
go 中变量的使用：
1. 内建变量类型
	* bool, string
	* (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
	* byte, rune  rune代表一字节, 32位
	* float32, float64, complex64, complex128  complex代表复数，实部和虚部分别为float类型
2. 强制类型转换
	类型转换是强制的，必须显示的转换类型
*/
package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// 函数外部变量不能使用 := 来定义变量，只能使用关键字定义
// a := 3 ❌
// 变量的作用域不是全局而是 package 内部
var globalA = 3
var globalB = true

// 使用括号来避免多个var出现
var (
	globalS = "global"
	globalN = 3
)

// 使用var关键字定义变量
func variableZeroValue() {
	// 变量声明后会有初值，int = 0，string = ""
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitalValue() {
	var a int = 3
	// var b, c int = 4    多个变量同时定义要都有合理的初值，并且定义后都要被使用
	var s string = "hello, go"

	fmt.Println("variableInitalValue: ", a, s)
	fmt.Printf("\n")
}

func variableNoType() {
	// 变量声明时可以根据值自己判断值的类型
	var a, b, s = 3, true, "string"
	fmt.Println("variableNoType: ", a, b, s)
	fmt.Printf("\n")
}

func variableShoter() {
	// 变量定义时可以省略 var 使用 := 来定义
	// 变量再次赋值时不能使用 :=
	// := 只能用来给声明变量初值
	a, b, s := 3, true, "shorter"
	b = false
	fmt.Println("variableShoter:", a, b, s)
	fmt.Printf("\n")
}

func euler() {
	e := cmplx.Pow(math.E, 1i*math.Pi) + 1
	fmt.Println("欧拉公式：", e)
	fmt.Printf("\n")
	// 运行结果： (0+1.2246467991473515e-16i)
	// 复数实部和虚部都是浮点数，所以不是确切的 0
}

func triangle() {
	var a, b int = 3, 4
	var c int
	// c = math.Sqrt(a*a + b*b) ❌ sqrt参数为 float64 类型， 返回为 float64 类型
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println("强制类型转换：", c)
	fmt.Printf("\n")
}

func main() {
	variableZeroValue()
	variableInitalValue()
	variableNoType()
	variableShoter()
	fmt.Println("global variables", globalA, globalB, globalS, globalN)
	fmt.Printf("\n")
	euler()
	triangle()
}
