package main

import (
	"fmt"
)

func main() {

	//string底层是一个byte数组，所以可以进行切片
	str := "hello@atguigu"
	//使用切片获得atguigu
	slice := str[6:]
	fmt.Println(slice)

}
