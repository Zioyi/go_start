package main

import "fmt"

func main() {
	map1 := make(map[string]interface{})

	map1["ids"] = []int{1, 2, 3}
	map1["name"] = "zioyi"

	fmt.Println(map1)

	ids := map1["ids"].([]int)

	fmt.Println(ids)

	name := map1["name"].(string)
	fmt.Println(name)

}
