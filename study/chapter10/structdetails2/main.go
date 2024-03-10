package main

import (
	"fmt"
)

type A struct {
	Num int
}
type B struct {
	Num int
}

type Monster struct {
	Name  string
	Age   int
	skill string
}

func main() {
	var a A
	var b B
	a = A(b) //b的实例可以强转为A类型，前提是结构体的字段要一致（名字、个数、类型）
	fmt.Println(a, b)

}
