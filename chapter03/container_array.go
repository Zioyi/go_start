package main

import (
	"fmt"
)

func traverseArray(a [3]int) {
	for i, v := range a {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}
}

func makeBoundary(msg string) {
	fmt.Println("xxxxxxxxxx", msg, "xxxxxxxxxx")
}

func main() {
	fmt.Println("10-04")

	var a [3]int
	fmt.Println(a[0])
	fmt.Println(a[len(a)-1])

	for i, v := range a {
		fmt.Printf("index: %d, value: %d\n", i, v)
	}

	for _, v := range a {
		fmt.Printf("value: %d\n", v)
	}
	makeBoundary()

	var q [3]int = [3]int{1, 2, 3}
	traverseArray(q)
	makeBoundary("")

	var r [3]int = [3]int{1, 2}
	traverseArray(r)
	makeBoundary("")

	s := [...]int{1, 2, 3}
	fmt.Printf("%T\n", s)
	makeBoundary("")

	x := [3]int{1, 2, 3}
	traverseArray(x)
	x = [3]int{1, 1, 1}
	traverseArray(x)
	// err: cannot use [4]int literal (type [4]int) as type [3]int in assignment
	// x = [4]int{1, 2, 3, 4}

	makeBoundary("比较两个数组是否相等")
	// 只有数组长度相等，类型相同才可以直接用 == 比较
	a1 := [2]int{1, 2}
	b1 := [...]int{1, 2}
	c1 := [2]int{1, 3}
	fmt.Println(a1 == b1, a1 == c1, b1 == c1)

	// d1 := [3]int{1, 2}
	//invalid operation: a1 == d1 (mismatched types [2]int and [3]int)
	// fmt.Println(a1 == d1)

}
