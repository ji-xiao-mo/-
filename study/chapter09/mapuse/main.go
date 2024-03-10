package main

import (
	"fmt"
)

func main() {
	//第一种使用方式
	var a map[string]string = make(map[string]string, 10)
	//使用map之前要先make，给map分配数据空间

	a["姓名1"] = "韩宗铮"
	a["姓名2"] = "武松"
	fmt.Println(a)

	//第二种使用方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "南京"
	fmt.Println(cities)
	//第三种使用
	heroes := map[string]string{
		"hero1": "鲁智深",
		"hero2": "李逵",
		"hero3": "武松",
	}
	fmt.Println(heroes)

	//案例演示 存放3个学生的信息，每个学生有name和sex的信息
	students := make(map[string]map[string]string)
	students["no1"] = make(map[string]string)
	students["no1"]["name"] = "tom"
	students["no1"]["sex"] = "男"

	fmt.Println(students)
}
