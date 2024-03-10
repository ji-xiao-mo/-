package main

import (
	"fmt"
)

type MethodUtils struct {
}

func (m MethodUtils) ju() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (M MethodUtils) juxing(m int, n int) {
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func (M MethodUtils) mianji(len float64, width float64) float64 {
	res := len * width
	return res
}

func main() {

	// 编写结构体(MethodUtils),编程一个方法，方法不需要参数，
	// 在方法中打印一个10*8的矩形,在main方法中调用该方法
	var M MethodUtils
	M.ju()
	M.juxing(5, 5)
	area := M.mianji(6.5, 5.555)
	fmt.Println(area)

}
