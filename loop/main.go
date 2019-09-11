package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
数字转换为二进制
for 语句可以没有初始值

strconv 模块用于字符串转换
*/
func convertToBinary(v int) string {
	result := ""
	if v < 0 {
		panic(fmt.Sprintln("Value must max then 0"))
	} else if v == 0 {
		result = "0"
	}
	for ; v > 0; v /= 2 {
		lbs := v % 2
		result = strconv.Itoa(lbs) + result
	}
	return result
}

/*
打印文件内容
for 语句可以没有初始值和递增表达式 相当于 while

os 提供平台无关的操作系统功能接口
bufio 缓存io
*/
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

/*
死循环
for 没有任何语句 相当于 while(true)
*/
func forever() {
	for {
		fmt.Println("forever")
	}
}

func main() {
	fmt.Println(
		convertToBinary(5),
		convertToBinary(13),
		convertToBinary(1024),
		convertToBinary(0),
	)
	printFile("abc.txt")
}
