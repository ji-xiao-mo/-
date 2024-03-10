package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i bool = false
	fmt.Println(i)
	fmt.Println("i占的空间=", unsafe.Sizeof(i))
}
