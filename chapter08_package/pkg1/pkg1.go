package pkg1

import (
	"fmt"
	"go_start/chapter08/pkg2"
)

func ExecPkg1() {
	fmt.Println("ExecPkg1")

	pkg2.ExecPkg2()
}

func init() {
	fmt.Println("pkg1 init")
}
