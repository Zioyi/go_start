package main

import (
	"fmt"
	"go_start/utils"
	"reflect"
)

type Command struct {
	Name    string
	Var     *int
	Commnet string
}

func newCommand(name string, varref *int, comment string) *Command {
	return &Command{
		Name:    name,
		Var:     varref,
		Commnet: comment,
	}
}

func main() {
	fmt.Println("Go语言实例化结构体——为结构体分配内存并初始化")

	utils.MakeBoundary("基本的实例化形式")
	type Point struct {
		X int
		Y int
	}
	// var ins T
	// T为结构体类型，ins为结构体的实例
	var p Point
	p.X = 100
	p.Y = 200
	fmt.Printf("type of (ins)p is: %v, value is %v\n", reflect.TypeOf(p), p)

	utils.MakeBoundary("创建指针类型的结构体")
	// ins := new(T)
	// ins的类型为*T
	p2 := new(Point)
	p2.X = 1
	fmt.Printf("type of (ins)p2 is: %v, value is %v\n", reflect.TypeOf(p2), p2)

	utils.MakeBoundary("取结构体的地址实例化")
	// ins := &T{}
	// ins的类型为*T
	p3 := &Point{
		X: 1,
		Y: -13,
	}
	fmt.Printf("type of (ins)p3 is: %v, value is %v\n", reflect.TypeOf(p3), p3)

	var version int = 1
	cmd := newCommand(
		"version",
		&version,
		"show version",
	)
	fmt.Println(cmd)

}
