package utils

import (
	"fmt"
)

// MakeBoundary :: make an obvious boundary
func MakeBoundary(msg string) {
	fmt.Printf("\n")
	fmt.Println("## new session ##")
	fmt.Println(msg)
}
