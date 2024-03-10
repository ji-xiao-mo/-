package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//有符号整数
	var i int8 = 127
	//var i int8 = 128  超出范围，报错
	fmt.Println("i=", i)

	//无符号整数uint 范围0～255
	var k uint8 = 255
	//var m uint8 = 299 超出范围
	fmt.Println("k=", k)

	//整形的使用细节
	var n1 = 100.0
	//查看某个变量的类型
	//fmt.printf()用于格式化输出
	fmt.Printf("n1的类型 %T \n", n1)
	//unsafe是个包，Sizeof是其中的函数可以返回变量大小
	fmt.Printf("n1占的空间大小 %d", unsafe.Sizeof(n1))
}
