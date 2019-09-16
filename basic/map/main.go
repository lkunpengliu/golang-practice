package main

import "fmt"

func main() {
	// map 的定义 map[K]V 方法一
	m := map[string]string{
		"name":    "apple",
		"country": "China",
		"quality": "good",
		"color":   "red",
	}
	// map 的定义 方法二
	m2 := make(map[string]int) // m2 == empty map
	// map 的定义 方法三
	var m3 map[string]string // m3 == nil
	fmt.Println("map ->", m)
	fmt.Println("map2 ->", m2)
	fmt.Println("map3 ->", m3)

	// map的遍历
	// 遍历是不会保证 map 每次的遍历顺序的
	for k, v := range m {
		fmt.Printf("key -> %s, value -> %s\n", k, v)
	}

	// 获取 map 中的值
	name := m["name"]
	fmt.Println("name ->", name)
	// 获取 map 中不存在的值
	emptyName := m["emptyname"]
	fmt.Println("emptyName ->", emptyName) // emptyName -> 会获取到空值，并不会报错
	// 获取 map 中的值的正确操作
	if name, ok := m["name"]; ok {
		fmt.Println("获取 map 中的值的正确操作 ->", name)
	} else {
		fmt.Println("Key does not exist")
	}

	// 删除 map 中的值
	color, ok := m["color"]
	fmt.Println("color ->", color)
	delete(m, "color")
	if color, ok = m["color"]; ok {
		fmt.Println("获取 map 中的值的正确操作 ->", color)
	} else {
		fmt.Println("Key does not exist")
	}
}
