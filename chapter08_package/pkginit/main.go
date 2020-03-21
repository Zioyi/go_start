package main

import (
	"fmt"
	"go_start/chapter08/pkg1"
)

func init() {
	fmt.Println("main init")
}

func main() {
	/*
		* pkg init() 初始化顺序
		* main <- pkg1 <- pkg2
		* so 1. pkg2.init() 2. pkg2.init() 3. main.init()

		>>> go run src/go_start/chapter08_package/pkginit/main.go
			pkg2 init
			pkg1 init
			main init
			ExecPkg1
			ExecPkg2
	*/
	pkg1.ExecPkg1()
}
