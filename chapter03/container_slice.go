package main

import (
	"fmt"

	"go_start/utils"
)

func main() {
	fmt.Println("10-04")

	var a = [3]int{1, 2, 3}
	fmt.Println(a, a[1:2])

	var a1 []int
	a1 = append(a1, 1)
	a1 = append(a1, 1, 2, 3)
	a1 = append(a1, []int{1, 2, 3}...)
	fmt.Println(a1)

	utils.MakeBoundary("切片在扩容时，容量的扩展规律是按容量的 2 倍数进行扩充，例如 1、2、4、8、16……，")
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len:%d cap:%d pointer:%p\n", len(numbers), cap(numbers), numbers)
	}

	utils.MakeBoundary("除了在切片尾部追加，我们还可以在切片开头添加元素（但性能会差很多）")
	var a2 = []int{1, 2, 3}
	a2 = append([]int{0}, a2...)
	fmt.Println(a2)
	a2 = append([]int{-3, -2, -1}, a2...)
	fmt.Println(a2)

	utils.MakeBoundary("因为append函数新切片的特性，所以切片也支持链式操作，我们可以将多个append操作组合起来，实现切片中间插入元素")
	var a3 = []int{1, 2, 3, 4, 6}
	// TODO: 有没有查找元素未知的函数
	a3 = append(a3[:4], append([]int{5}, a3[4:]...)...)
	fmt.Println(a3)

}
