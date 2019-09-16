package main

import "fmt"

func printSlice(arr []int) {
	fmt.Printf("s=%v, len=%d, cap=%d\n", arr, len(arr), cap(arr))
}

func sliceops() {
	fmt.Println("Creating slice......")
	// 定义一个 slice, slice的初始值为 nil
	var s []int
	fmt.Println("s ->", s)
	fmt.Println(s == nil)

	// append后如果原数组不够大，新的数组cap会是原数组cap的2倍
	for i := 0; i < 8; i++ {
		printSlice(s)
		s = append(s, i)
	}

	fmt.Println("s ->", s)

	// 使用 make 方法创建 slice
	// make(Type, len, cap)  type 为 slice 存储的数据类型， len 为数组长度，cap 为容量大小
	s2 := make([]int, 15)
	s3 := make([]int, 16, 32)
	printSlice(s2)
	printSlice(s3)

	fmt.Println("Copying slice......")
	copy(s2, s)
	printSlice(s2)

	fmt.Println("Deleting elements from slice......")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
}
