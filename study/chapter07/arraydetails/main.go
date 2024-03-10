package main

import (
	"fmt"
)

func main() {
	//数组一旦声明/定义了，长度固定不变，不能动态变化
	var arr01 [3]int
	arr01[0] = 1
	arr01[1] = 30
	arr01[2] = 1

	fmt.Println(arr01)
}
