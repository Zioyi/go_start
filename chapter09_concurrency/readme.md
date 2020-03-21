## Goroutine 介绍
goroutine是比线程还要轻量级的抽象，由go的runtime负责调度，开发者只管声明，我们可以在一个进程里声明多个协程来*并发*执行任务。

goroutine的用法
```
// go关键字放在方法调用前新建一个 goroutine 并执行方法体
go GetThingDone(param1, param2)

// 新建一个匿名方法并执行
go func(param1, param2) {
}(val1, val2)

// 直接新建一个goroutine 并在 goroutine 中执行代码块
go {
    //do something
}

```
---

## channel
channel是一种 goroutine 间的通信方式。一个channel只能传递一种类型的值。

channel必须是用make创建，例：
```
ci := make(chan int)
cs := make(chan string)
cf := make(chan intergace{})
```

### 单向通道的声明方式
```
var 通道实例 chan<- 元素类型 //只能发送通道
var 通道实例 <-chan 元素类型 //只能接收通道
```

示例：
```
ch := make(chan int)
// 声明一个只能发送的通道类型，并赋值为ch
var chSendOnly chan<- int = ch
// 声明一个只能接收的通道类型，并赋值为ch
var chRecvOnlu <-chan int = ch


```