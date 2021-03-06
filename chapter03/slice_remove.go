package main

import (
	"fmt"
	"go_start/utils"
)

func main() {
	fmt.Println("Go语言从切片中删除元素")

	utils.MakeBoundary("Go语言并没有对删除切片元素提供专用的语法或者接口，需要使用切片本身的特性来删除元素。根据要删除元素的位置有三种情况：从开头位置删除，从中间位置删除，从尾部删除。其中删除切片尾部的元素最快：")
	var a = []int{1, 2, 3}
	a = a[:len(a)-1] // 删除尾部1个元素
	fmt.Println(a)
	var N = 1
	a = a[:len(a)-N] // 删除尾部N个元素
	fmt.Println(a)

	utils.MakeBoundary("删除开头的元素可以直接移动数据指针：")
	a = []int{1, 2, 3}
	a = a[1:]
	fmt.Println(a)

	utils.MakeBoundary("删除开头的元素也可以不移动数据指针，但是将后面的数据向开头移动。可以用 append 原地完成（所谓原地完成是指在原有的切片数据对应的内存区间内完成，不会导致内存空间结构的变化）：")
	a = []int{1, 2, 3}
	a = append(a[:0], a[1:]...)
	fmt.Println(a)

	utils.MakeBoundary("也可以用 copy 完成删除开头的元素：")
	a = []int{1, 2, 3}
	// rv := copy(a, a[1:])
	// fmt.Println(rv)
	// fmt.Println(a)
	a = a[:copy(a, a[1:])]
	fmt.Println(a)

	utils.MakeBoundary("对于删除中间的元素，需要对剩余的元素进行一次整体挪动，同样可以用 append 或 copy 原地完成：")
	a = []int{1, 2, 3}
	i := 1
	a = append(a[:i], a[i+1:]...)
	fmt.Println(a)

	N = 2
	a = []int{1, 2, 3, 4}
	i = 1
	a = append(a[:i], a[i+N]) // 删除中间N个元素
	fmt.Println(a)

	a = []int{1, 2, 3}
	i = 1
	new_a := a[:i+copy(a[i:], a[i+1:])]
	fmt.Println(a)
	fmt.Println(new_a)

	a = []int{1, 2, 3, 4}
	i = 1
	N = 2
	a = a[:i+copy(a[i:], a[i+N:])]
	fmt.Println(a)
}
