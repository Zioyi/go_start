package main

import (
	"fmt"
	"go_start/utils"
)

func main() {
	fmt.Println("Go语言copy()：切片复制（切片拷贝）")

	utils.MakeBoundary("Go语言的内置函数 copy() 可以将一个数组切片复制到另一个数组切片。如果加入的两个数组切片不一样大，就会按其中较小的那个数组切片的元素个数进行复制。")
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	copy(slice2, slice1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	slice2 = []int{5, 4, 3}
	copy(slice1, slice2)
	fmt.Println(slice1)
	fmt.Println(slice2)

	utils.MakeBoundary("通过代码演示对切片的引用和复制操作后对切片元素的影响")
	const elementCount = 1000
	srcData := make([]int, elementCount)
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}

	refData := srcData

	copyData := make([]int, elementCount)
	copy(copyData, srcData)

	srcData[0] = 999

	fmt.Println(refData[0])
	fmt.Println(copyData[0], copyData[elementCount-1])

	copy(copyData, srcData[4:6])
	// fmt.Println(copyData)
	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}
