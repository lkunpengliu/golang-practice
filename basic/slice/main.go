/*
切片
*/
package main

import "fmt"

func updateArray(arr []int) {
	arr[0] = 10
}

/*
向slice添加元素
	* 添加元素时如果超越cap，系统会重新分配更大的底层数组
	* 由于值传递的关系，必须接收 append 的返回值
	* s = append(s, val)
*/
func appendSlice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	s1 := arr[2:6]
	s2 := s1[3:5]

	fmt.Println("s1 ->", s1)
	fmt.Println("s2 ->", s2)

	fmt.Printf("s2=%v, len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))

	s3 := append(s2, 10)
	// s4 和 s5 不再是 arr 的view，而是指向了一个系统分配的新的更大的数组
	s4 := append(s3, 11)
	s5 := append(s4, 12)

	fmt.Println("s3, s4, s5", s3, s4, s5)
	fmt.Println("arr ->", arr)
}

func main() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	// slice 本身没有数据，是对底层array的一个view
	s1 := arr[2:6]
	s2 := arr[:6]
	s3 := arr[2:]
	s4 := arr[:]

	fmt.Println(
		s1,
		s2,
		s3,
		s4,
	)
	// 由于 slice 是对底层 array 的一个view，所以对 slice 的修改会反应到底层 array
	fmt.Println("after updateArray(s1)")
	updateArray(s1)
	fmt.Println("s1 ->", s1)
	fmt.Println("arr ->", arr)

	fmt.Println("Reslice")
	fmt.Println("before reslice s3 ->", s3)
	s3 = s3[:5]
	fmt.Println("Reslice s3 ->", s3)

	/*
		slice 扩展
			* slice 可以向后扩展，不可以向前扩展
			* s[i]不可以超越 len(s)，向后扩展不可以超越底层数组 cap(s)
		slice 源码 https://golang.org/src/runtime/slice.go
		可参考文章 https://juejin.im/post/5c8f30f3f265da612b1aad29
	*/
	arr2 := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	sl1 := arr2[2:6]
	sl2 := sl1[3:5]

	fmt.Println("slice 扩展")
	fmt.Println("sl1 ->", sl1)
	fmt.Println("sl2 ->", sl2) // sl1[3] sl1[4]
	// fmt.Println("sl1[4] ->", sl1[4]) //panic: runtime error: index out of range [4] with length 4

	fmt.Printf("sl1=%v, len(sl1)=%d, cap(sl1)=%d\n", sl1, len(sl1), cap(sl1))
	fmt.Printf("sl2=%v, len(sl2)=%d, cap(sl2)=%d\n", sl2, len(sl2), cap(sl2))

	fmt.Println("slice append")
	appendSlice()

	sliceops()
}
