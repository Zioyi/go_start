package main

import (
	"fmt"
	"reflect"
)

//在结构体成员嵌入时使用别名

type Brand struct {
}

func (t Brand) Show() {
	fmt.Printf("hello world!\n")
}

type FakeBrand = Brand

type Vehicle struct {
	FakeBrand
	Brand
}

func main() {
	var a Vehicle

	a.FakeBrand.Show()
	a.FakeBrand.Show()

	// 取a的类型反射对象
	ta := reflect.TypeOf(a)

	for i := 0; i < ta.NumField(); i++ {

		// a的成员信息
		f := ta.Field(i)

		fmt.Printf("FieldName: %v, FieldType: %v\n", f.Name, f.Type.Name())
	}

}
