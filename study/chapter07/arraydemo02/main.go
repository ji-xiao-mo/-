package main

import (
	"fmt"
)

func main() {

	var intArr [3]int
	//数组的默认元素已经有初始值
	fmt.Println(intArr)
	fmt.Printf("%p %p\n", &intArr, &intArr[1])

	//案例 从终端循环输入5个成绩，保存到float64数组当中，并输出
	var score [5]float64

	for i := 0; i < len(score); i++ {
		fmt.Printf("请输入第%d个成绩", i+1)
		fmt.Scanln(&score[i])
	}
	fmt.Println(score)

}
