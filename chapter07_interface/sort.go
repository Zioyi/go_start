package main

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	names := MyStringList{
		"aaa",
		"ccc",
		"bbb",
	}

	sort.Sort(names)

	for _, v := range names {
		fmt.Printf("%s\n", v)
	}
}
