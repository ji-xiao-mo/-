package main

import (
	"fmt"
)

func main() {

	//map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 55
	map1[2] = 66
	map1[7] = 32
	fmt.Println(map1)

}
