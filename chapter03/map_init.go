package main

import (
	"fmt"
	"go_start/utils"
)

func main() {
	fmt.Println("Go语言map（Go语言映射）")

	utils.MakeBoundary("key可以是任意支持==或!=操作符的类型，比如：string、int、float。所以数组、切片、和结构体不能作为key，但指针和接口类型可以。如果要用结构体作为key可以提款Key()和Hash()方法，这样可以通过结构体的域计算出唯一的数字或字符串的key。")
	var mapLit map[string]int
	var mapAssigned map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit
	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is :%d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is :%f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is :%d\n", mapAssigned["two"])
	fmt.Printf("Map literal at \"two\" is :%d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is :%d\n", mapLit["ten"])

	utils.MakeBoundary("不要使用 new，永远用 make 来构造 map")
	mapCreatedByNew := new(map[string]float32)
	fmt.Println(mapCreatedByNew)
	// invalid operation: mapCreatedByNew["key1"] (type *map[string]float32 does not support indexing)
	// mapCreatedByNew["key1"] = 1.0

	utils.MakeBoundary("map容量 当 map 增长到容量上限的时候，如果再增加新的 key-value 对，map 的大小会自动加 1。所以出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。")
	map2 := make(map[string]float32, 1)
	fmt.Printf("Map2 cap is :%d\n", len(map2))

	utils.MakeBoundary("切片作为map的值 1.值切片长度固定")
	map3 := make(map[int][2]int)
	map3[0] = [2]int{}
	map3[1] = [2]int{1}
	map3[2] = [2]int{1, 1}
	fmt.Println(map3)

	utils.MakeBoundary("切片作为map的值 2.值切片长度不固定")
	map4 := make(map[int][]int)
	map4[0] = []int{}
	map4[1] = []int{1}
	map4[2] = []int{1, 1}
	map4[100] = append([]int{1}, 99)
	fmt.Println(map4)
}
