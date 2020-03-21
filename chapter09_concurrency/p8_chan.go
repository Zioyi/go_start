package main

import "fmt"

func main() {
	// 构建一个通道
	ch := make(chan int)

	go func() {
		fmt.Println("start goroutine")

		ch <- 0

		fmt.Println("exit goroutine")
	}()

	fmt.Println("wait goroutine")

	<-ch

	fmt.Println("all done")
}
