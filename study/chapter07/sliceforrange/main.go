package main

import (
	"fmt"
)

func main() {

	//使用常规的for循环的方式
	var arr [5]int = [...]int{2, 56, 76, 45, 76}
	slice := arr[1:4]
	fmt.Println(slice)
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//使用for--range方式遍历数组
	for i, v := range slice {
		fmt.Println(i, v)
	}

	slice2 := slice[1:2]
	fmt.Println(slice2)

	//append内置函数，可以对切片进行动态追加
	var slice3 []int = []int{10, 21, 43}
	//通过append来对slice3进行追加
	slice3 = append(slice3, 400)
	//可以直接追加切片
	slice3 = append(slice3, slice2...)
	fmt.Println(slice3)

	//拷贝操作
	var slice4 []int = []int{3, 4, 7}
	var slice5 []int = make([]int, 5, 10)
	copy(slice5, slice4)
	fmt.Println("slice5", slice5)
}
