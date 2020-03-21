package main

import (
	"fmt"
	"go_start/utils"
)

func main() {
	fmt.Println("Go语言初始化结构体的成员变量")

	utils.MakeBoundary("使用“键值对”初始化结构体")
	// ins := 结构体类型名{
	// 	字段1: 字段1的值,
	// 	字段2: 字段2的值,
	// 	…
	// }
	// 下面是对各个部分的说明：
	// 结构体类型：定义结构体时的类型名称。
	// 字段1、字段2：结构体的成员字段名。结构体类型名的字段初始化列表中，字段名只能出现一次。
	// 字段1 的值、字段2 的值：结构体成员字段的初始值。
	type People struct {
		name  string
		child *People
	}

	relation := People{
		name: "爷爷",
		child: &People{
			name: "爸爸",
			child: &People{
				name: "我",
			},
		},
	}
	fmt.Println(relation)
	fmt.Println(relation.child)
	fmt.Println(relation.child.child)

	utils.MakeBoundary("使用多个值的列表初始化结构体")
	// ins := 结构体类型名{
	// 	字段1的值,
	// 	字段2的值,
	// 	…
	// }
	// 使用这种格式初始化时，需要注意：
	// 必须初始化结构体的所有字段。
	// 每一个初始值的填充顺序必须与字段在结构体中的声明顺序一致。
	// 键值对与值列表的初始化形式不能混用。
	type Address struct {
		Province    string
		City        string
		ZipCode     int
		PhoneNumber string
	}

	addr := Address{
		"四川",
		"成都",
		610000,
		"0",
	}

	fmt.Println(addr)

	utils.MakeBoundary("初始化匿名结构体")
	// ins := struct {
	// 	// 匿名结构体字段定义
	// 	字段1 字段类型1
	// 	字段2 字段类型2
	// 	…
	// }{
	// 	// 字段值初始化
	// 	初始化字段1: 字段1的值,
	// 	初始化字段2: 字段2的值,
	// 	…
	// }
	msg := &struct {
		id   int
		data string
	}{
		1024,
		"hello",
	}
	printlnMsgType(msg)
}

func printlnMsgType(msg *struct {
	id   int
	data string
}) {
	fmt.Printf("%T\n%v\n", msg, msg)
}
