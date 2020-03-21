package main

import "fmt"

func main1() {
	fmt.Println("Hello World")
	var a string = "Runoob"
	fmt.Println(a)

	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	intVal := 1
	fmt.Println(intVal)

	// for true {
	// 	fmt.Printf("这是无限循环。\n")
	// }
}

// func
func max(num1 int, num2 *int) int {
	var result int

	*num2 = 0
	if num1 > *num2 {
		result = num1
	} else {
		result = *num2
	}
	return result
}

func main() {
	var a int = 100
	var b int = 200
	var ret int

	ret = max(a, &b)

	fmt.Printf("最大值是：%d\n", ret)
	fmt.Printf("b是：%d\n", b)
}
