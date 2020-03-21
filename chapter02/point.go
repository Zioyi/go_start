package main

import (
	"flag"
	"fmt"
)

func swap(a, b *int32) (*int32, *int32) {
	fmt.Printf("%x %d\n", a, *a)
	fmt.Printf("%x %d\n", b, *b)
	a, b = b, a

	return a, b

}

var mode = flag.String("mode", "", "process mode")

func main() {
	//从指针获取指针指向的值
	var house = "Malibu Point 10880, 90265"

	ptr := &house
	fmt.Printf("ptr type: %T\n", ptr)
	fmt.Printf("address: %p\n", ptr)

	value := *ptr
	fmt.Printf("value: %s\n", value)

	//使用指针修改值
	var x int32 = 1
	var y int32 = 2
	new_x, new_y := swap(&x, &y)
	fmt.Println(x, y)
	fmt.Println(*new_x, *new_y)

	//使用指针变量获取命令行的输入信息
	flag.Parse()
	fmt.Println(*mode)
}
