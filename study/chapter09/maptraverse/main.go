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
	//使用for range遍历
	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	students := make(map[string]map[string]string)
	students["no1"] = make(map[string]string)
	students["no1"]["name"] = "tom"
	students["no1"]["sex"] = "男"

	fmt.Println(students)
	fmt.Println("student有", len(students), "对")
	for k1, v1 := range students {
		fmt.Println(k1)
		for k2, v2 := range v1 {
			fmt.Println(k2, v2)
		}
	}
}
