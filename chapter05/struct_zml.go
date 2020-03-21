package main

import (
	"fmt"
	"go_start/chapter05/zml"
	"go_start/utils"
)

func main() {
	fmt.Println("Go语言结构体字面量")

	utils.MakeBoundary("方式一：按顺序传值")
	type Point struct{ X, Y int }
	p := Point{1, 2}

	fmt.Println(p)

	utils.MakeBoundary("方式二：指定名称和值来初始化")
	p1 := Point{
		X: 11,
	}
	fmt.Println(p1)

	pp2 := zml.PointF{11111, 2}
	fmt.Println(pp2)

	pp2 = zml.PointF{X: 1231412}
	fmt.Println(pp2)
}
