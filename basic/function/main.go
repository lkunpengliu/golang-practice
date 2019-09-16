package main

import "fmt"

/*
返回值类型在最后
可返回多个值
*/
func eval(a, b int, op string) (r int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("wrong operation %s", op)
	}
}

/*
函数作为参数
*/
func apply(op func(int, int) int, a, b int) int {
	return op(a, b)
}

func div(a, b int) int {
	return a / b
}

/*
没有默认参数，可以使用可选参数
*/
func sum(numbers ...int) int {
	sum := 0
	for i := range numbers {
		sum += numbers[i]
	}
	return sum
}

func main() {
	if res, err := eval(3, 4, "+"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(res)
	}

	if res, err := eval(3, 4, "x"); err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(res)
	}

	fmt.Println(apply(div, 10, 3))

	fmt.Println(
		sum(1, 2, 3, 4, 5))
}
