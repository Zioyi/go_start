package main

import (
	"fmt"
	"strings"
)

func main() {
	var str = "C语言中文网\nGo语言教程"
	fmt.Println(str)
	for _, s := range str {
		fmt.Printf("%c\n", s)
	}

	fmt.Println("---")
	tracer := "死神来了， 死神bye bye"
	comma := strings.Index(tracer, "， ")
	fmt.Println(tracer[comma:])
	pos := strings.Index(tracer[comma:], "死神")
	fmt.Println(comma, pos, tracer[comma+pos:])

	fmt.Println("---")
	angel := "Heros never die"
	angelBytes := []byte(angel)
	for i := 5; i <= 10; i++ {
		angelBytes[i] = ' '
	}
	fmt.Println(string(angelBytes))

}
