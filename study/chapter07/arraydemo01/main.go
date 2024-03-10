package main

import (
	"fmt"
)

func main() {

	//定义一个数组
	var hens [6]float64
	//给数组元素赋值
	for i := 0; i < 6; i++ {
		hens[i] = 1.5
	}
	//数组循环求和
	res := 0.0
	for i := 0; i < len(hens); i++ {
		res += hens[i]
	}
	//求平均体重
	fmt.Println(res / float64(len(hens)))
}
