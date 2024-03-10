package main

import (
	"fmt"
)

func main() {

	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "南京"
	fmt.Println(cities)

	//因为no3已经存在，所以就是修改
	cities["no3"] = "南京~"
	fmt.Println(cities)

	//演示删除
	delete(cities, "no1")
	fmt.Println(cities)
	//删除了不存在的映射，不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	//map的查找
	value, key := cities["no2"]
	if key {
		fmt.Println("存在no2的key, 值为", value)
	} else {
		fmt.Println("不存在no2的key")
	}

	//全部删除
	cities = make(map[string]string)
	fmt.Println(cities)

}
