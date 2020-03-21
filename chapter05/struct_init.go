package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go语言结构体定义")

	type Point struct {
		X int
		Y int
	}

	var p Point
	fmt.Println(p)
	fmt.Printf("StructP X is :%d, Y is :%d\n", p.X, p.Y)

}
