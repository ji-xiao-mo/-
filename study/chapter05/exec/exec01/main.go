package main

import (
	"fmt"
)

func main() {
	//第一个作业，声明两个变量 i和j 然后判断两数之和与50的大小关系，
	//大于50的话，输出hello world
	var i int32
	var j int32
	fmt.Println("请输入i的值：")
	fmt.Scanln(&i)
	fmt.Println("请输入j的值：")
	fmt.Scanln(&j)

	var m int32 = i + j

	if m > 50 {
		fmt.Println("hello,world")
	} else {
		fmt.Println("数值小于50")
	}
	//第二个作业声明两个浮点数，若第一个数大于，第二个数小于，则输出
	var n1 float64 = 120.3
	var n2 float64 = 21.53

	if n1 > 10.0 && n2 < 90.0 {
		fmt.Println("两数之和为：", n1+n2)
	}
	//第三个作业，定义两个int32型的数据，
	//判断两者的和，是否能被3整除又能被5整除，打印提示信息
	var n3 int32 = 21
	var n4 int32 = 2
	var n5 = n3 + n4
	if n5%3 == 0 && n5%5 == 0 {
		fmt.Println("可同时被5和3整除")
	} else {
		fmt.Println("不能同时被5和3整除")
	}

	//第四个作业，判断闰年
	var year int32
	fmt.Println("请输入年份：")
	fmt.Scanln(&year)
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println(year, "是闰年")
	} else {
		fmt.Println(year, "不是闰年")
	}

}
