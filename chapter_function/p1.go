package main

import "fmt"

func typedTwoValues() (int, int) {
	return 1, 2
}

func main() {
	a, b := typedTwoValues()
	fmt.Println(a, b)
}
