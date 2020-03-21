package main

import (
	"fmt"
	"go_start/utils"
)

func main() {
	fmt.Println("Go语言range关键字：循环迭代切片")

	utils.MakeBoundary("既然切片是一个集合，那么就可以迭代其中的元素。Go语言有个特殊的关键字 range，它可以配合关键字 for 来迭代切片里的元素，如下所示：")
	slice := []int{10, 20, 30, 40}
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	utils.MakeBoundary("当迭代切片时，关键字 range 会返回两个值。第一个值是当前迭代到的索引位置，第二个值是该位置对应元素值的一份副本,需要强调的是，range 创建了每个元素的副本，而不是直接返回对该元素的引用")
	for index, value := range slice {
		fmt.Printf("Value: %d Value-Addr: %X ElemAddr: %X\n",
			value, &value, &slice[index])
	}

}
