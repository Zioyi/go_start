package main

import "fmt"

func printer(c chan int) {

	for {
		data := <-c

		if data == 0 {
			break
		}

		fmt.Println(data)
	}

	// 通知main已经结束循环（我搞定了！）
	c <- 0
}

func main() {
	c := make(chan int)

	// 这里声明了3个 goroutine，它们会争抢chan里的数据
	go printer(c)
	go printer(c)
	go printer(c)

	for i := 1; i <= 10; i++ {
		c <- i
	}

	// 通知并发的printer结束循环（没数据啦！）
	c <- 0

	// 等待printer结束（搞定喊我！）
	<-c
}
