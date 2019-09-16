// 操作字符串的包 strings
package main

import (
	"fmt"
)

func main() {
	s := "Hello,欢迎来到我的世界!" // UTF-8编码
	// 字符串长度
	l := len(s)
	fmt.Println("字符串长度l ->", l)            // 字符串长度l -> 31 中文字符以16进制编码，每一个中文有3个字节
	fmt.Printf("正常输出字符串s -> %s\n", s)      // 正常输出字符串 -> Hello,欢迎来到我的世界!
	fmt.Printf("16进制方式输出字符串s -> % X\n", s) // 48 65 6C 6C 6F 2C E6 AC A2 E8 BF 8E E6 9D A5 E5 88 B0 E6 88 91 E7 9A 84 E4 B8 96 E7 95 8C 21

	for i, ch := range s { // ch is a rune
		fmt.Printf("(%d, %X) ", i, ch)
	}
	fmt.Println()

	// rune 相当于 go 的 char
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c)", i, ch)
	}
	fmt.Println()
}
