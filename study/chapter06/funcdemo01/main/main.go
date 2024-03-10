package main

import (
	//引到包所在的文件夹
	"GOwork/study/chapter06/funcdemo01/utils"
	"fmt"
)

func main() {

	var n1 float64
	var n2 float64
	var operator string

	fmt.Println("请输入n1")
	fmt.Scanln(&n1)

	fmt.Println("请输入n2")
	fmt.Scanln(&n2)

	fmt.Println("请输入运算符")
	fmt.Scan(&operator)

	fmt.Println("结果为：dwdw", utils.Cal(n1, n2, operator))

}
