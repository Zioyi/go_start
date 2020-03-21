package main

import "fmt"

// Dictionary 字典结构
type Dictionary struct {
	data map[interface{}]interface{} // 键值都是interface{}类型
}

// Get 根据键获取值
func (d *Dictionary) Get(key interface{}) interface{} {
	return d.data[key]
}

// Set 设置键值
func (d *Dictionary) Set(key interface{}, value interface{}) {
	d.data[key] = value
}

// Visit 遍历所有的值，如果回调函数为false，停止遍历
func (d *Dictionary) Visit(callback func(k, v interface{}) bool) {

	if callback == nil {
		return
	}

	for k, v := range d.data {
		if !callback(k, v) {
			return
		}
	}
}

// Clear 清空所有数据
func (d *Dictionary) Clear() {
	d.data = make(map[interface{}]interface{})
}

// NewDictionary 创建一个字典
func NewDictionary() *Dictionary {
	d := &Dictionary{}

	// 初始化map
	d.Clear()
	return d
}

func main() {
	dict := NewDictionary()

	dict.Set("My Factory", 60)
	dict.Set("Terra Craft", 36)
	dict.Set("Don't Hungry", 24)

	favorite := dict.Get("Terra Craft")
	fmt.Println("favorite:", favorite)

	dict.Visit(func(key, value interface{}) bool {

		//将值转为int类型，并判断是否大于40
		if value.(int) > 40 {

			fmt.Println(key, "is expensive")
			return true
		}

		fmt.Println(key, "is cheap")
		return true
	})
}
