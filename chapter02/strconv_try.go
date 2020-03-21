package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("a" + strconv.Itoa(32))

	i, _ := strconv.Atoi("3")
	fmt.Println(3 + i)

	i, err := strconv.Atoi("A")
	if err != nil {
		fmt.Println("converted failed")
		fmt.Println(err)
	}
}
