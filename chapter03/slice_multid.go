package main

import (
	"fmt"
)

func main() {
	fmt.Println("Go语言多维切片简介")

	slice := [][]int{{10}, {100, 200}}
	fmt.Println(slice)

	slice[0] = append(slice[0], 20)
	fmt.Println(slice)
}
