package main

import (
	"fmt"
)

func main() {

	//演示make方法使用
	var slice []float64 = make([]float64, 5, 10)
	//对于切片，必须make后才能使用
	slice[0] = 10
	fmt.Println(slice)
	fmt.Println("slice的大小", len(slice))
	fmt.Println("slice的容量", cap(slice))

	//方式三
	//定义切片，直接指定具体数组
	var strslice []string = []string{"tom", "fwe", "jai"}
	fmt.Println(strslice)
	fmt.Println("strslice的大小", len(strslice))
	fmt.Println("strslice的容量", cap(strslice))

}
