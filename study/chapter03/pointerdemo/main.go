package main

import (
	"fmt"
)

// 演示golong的指针类型
func main() {

	//基本数据类型的内存布局
	var i int = 10
	//查询i的地址
	fmt.Println(&i)

	var ptr *int = &i
	fmt.Println(ptr)
	fmt.Println(&ptr)
	fmt.Println(*ptr)

}
