package main

import "fmt"

type A struct {
	title string
}

func (a *A) GetTitle() string {
	return a.title
}

func (a *A) GetTitle2() string {
	return a.title
}

type B struct {
	title string
	A
}

func (b *B) GetTitle() string {
	return b.title
}

func main() {

	b := new(B)

	fmt.Printf("%v\n", b)
	b.title = "b_title"
	fmt.Printf("%v\n", b)
	b.A.title = "a_title"
	fmt.Printf("%v\n", b)

	fmt.Println(b.GetTitle())
	fmt.Println(b.A.GetTitle())
	fmt.Println(b.GetTitle2())
}
