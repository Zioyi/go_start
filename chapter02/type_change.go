package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("int8 range:\t", math.MinInt8, "\t\t\t", math.MaxInt8)
	fmt.Println("int16 range:\t", math.MinInt16, "\t\t", math.MinInt16)
	fmt.Println("int32 range:\t", math.MinInt32, "\t\t", math.MaxInt32)
	fmt.Println("int64 range:\t", math.MinInt64, "\t", math.MaxInt64)
	//9,223,372,036,854,775,807

	var a int32 = 1047483647
	fmt.Printf("int32: 0x%x %d\n", a, a)

	b := int16(a)
	fmt.Printf("int16 0x%x %d\n", b, b)

	var c float32 = math.Pi
	fmt.Println(c, int(c))

}
