package main

import (
	"fmt"
)

func fbn(n int) []uint64 {
	//声明切片，大小为n
	fbnSlice := make([]uint64, n)
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice

}

func main() {

	/*
		思路：
		1、声明一个函数 fbn(n int)([]uint64)
		2、编程fbn(n int) 进行for循环存放斐波那契数列
	*/
	fbnslice := fbn(10)
	fmt.Println(fbnslice)

}
